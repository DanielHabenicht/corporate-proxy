package cmd

import (
	"fmt"
	// "os"

	"github.com/spf13/cobra"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("on called")
	//Checking that an environment variable is present or not.
		// redisHost, ok := os.LookupEnv("REDIS_HOST")
	// if !ok {
	//   fmt.Println("REDIS_HOST is not present")
	// } else {
	//   fmt.Printf("Redis Host: %s\n", redisHost)
	// }

	// os.Setenv("HTTP_PROXY", "1")
	// os.Setenv("HTTPS_PROXY", "1")
    // fmt.Println("HTTP_PROXY", os.Getenv("HTTP_PROXY"))
    // fmt.Println("HTTPS_PROXY", os.Getenv("HTTPS_PROXY"))

	// os.Unsetenv("SITE_TITLE")
	},
}

func init() {
	rootCmd.AddCommand(onCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
