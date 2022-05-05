package cmd

import (
	"gophkeeper/pkg/logger"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorization",
	Long:  `Choose one of the command to do with authorization`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.CheckErr(cmd.Help())
	},
}

var authRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register on the remote server",
	Long:  `Allows you to register on the remote server`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login on the remote server",
	Long:  `Allows you to login on the remote server`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var authForgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Forget current authorization",
	Long:  `Allows you to forget current authorization`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(authRegisterCmd)
	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authForgetCmd)
}
