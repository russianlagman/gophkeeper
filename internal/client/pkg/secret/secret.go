package secret

import (
	"encoding/json"
	"fmt"
)

var (
	_ Secret = (*Card)(nil)
	_ Secret = (*LoginPassword)(nil)
	_ Secret = (*Raw)(nil)
)

const (
	TypeRaw           = "raw"
	TypeLoginPassword = "lp"
	TypeCard          = "card"
)

type Secret interface {
	Type() string
	Encode() ([]byte, error)
	Decode([]byte) error
	Print() string
}

func Read(t string, data []byte) (Secret, error) {
	var v Secret

	switch t {
	case TypeCard:
		v = &Card{}
	case TypeLoginPassword:
		v = &LoginPassword{}
	case TypeRaw:
		fallthrough
	default:
		v = &Raw{}
	}

	if err := v.Decode(data); err != nil {
		return nil, err
	}

	return v, nil
}

type Card struct {
	Number  string `json:"number"`
	Expires string `json:"expires"`
	CVV     string `json:"cvv"`
	Holder  string `json:"holder"`
}

func (s *Card) Type() string {
	return "card"
}

func (s *Card) Encode() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Card) Decode(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}

func (s *Card) Print() string {
	return fmt.Sprint(s)
}

type LoginPassword struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (s *LoginPassword) Type() string {
	return "lp"
}

func (s *LoginPassword) Encode() ([]byte, error) {
	return json.Marshal(s)
}

func (s *LoginPassword) Decode(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}

func (s *LoginPassword) Print() string {
	return fmt.Sprint(s)
}

type Raw []byte

func (s *Raw) Type() string {
	return "raw"
}

func (s *Raw) Encode() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Raw) Decode(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}

func (s *Raw) Print() string {
	return fmt.Sprint([]byte(*s))
}
