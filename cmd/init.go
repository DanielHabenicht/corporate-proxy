package cmd

import (
	"fmt"
	"corporate-proxy-cli/cmd/utils"

	"github.com/spf13/cobra"
)

var HTTP_Proxy string
// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		utils.Command()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.PersistentFlags().StringVarP(&HTTP_Proxy, "http_proxy", "", "", "HTTP Proxy (required)")
	rootCmd.MarkPersistentFlagRequired("http_proxy")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
