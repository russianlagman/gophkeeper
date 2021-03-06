package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gophkeeper/internal/server/config"
	"gophkeeper/internal/server/grpcservice"
	"gophkeeper/internal/server/migrate"
	"gophkeeper/internal/server/storage/postgres"
	"gophkeeper/pkg/grpcserver"
	"gophkeeper/pkg/logger"
	"gophkeeper/pkg/token"
)

type App struct {
	config config.Config
	logger logger.Logger
	stop   chan struct{}
	server *grpcserver.Server
}

func New(cfg config.Config) (*App, error) {
	l := *logger.Global()

	l.Debug().Str("dsn", cfg.DB.DSN).Msg("Connecting to db")

	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("db open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping: %w", err)
	}
	if err := migrate.Up(db); err != nil {
		return nil, fmt.Errorf("migrate up: %w", err)
	}

	tm, err := token.NewJWT(cfg.Security.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("token manager: %w", err)
	}

	users, err := postgres.NewUserRepository(db)
	if err != nil {
		return nil, fmt.Errorf("user repository: %w", err)
	}

	secrets, err := postgres.NewSecretRepository(db)
	if err != nil {
		return nil, fmt.Errorf("user repository: %w", err)
	}

	as := grpcservice.NewUser(users, tm)
	ks := grpcservice.NewKeeper(secrets)

	s := grpcserver.New(
		grpcserver.WithListenAddr(cfg.GRPC.ListenAddr),
		grpcserver.WithServices(as, ks),
		grpcserver.WithUnaryInterceptors(grpcservice.BuildUnaryInterceptors()...),
		grpcserver.WithAuthFunc(grpcservice.BuildAuthFunc(tm)),
	)

	if err := s.Start(); err != nil {
		return nil, fmt.Errorf("grpc: %w", err)
	}

	a := &App{
		config: cfg,
		logger: l,
		stop:   make(chan struct{}),
		server: s,
	}

	return a, nil
}

func (a *App) Stop() {
	close(a.stop)
	a.server.Stop()
}
