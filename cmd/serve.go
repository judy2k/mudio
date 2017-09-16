package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the mudio web service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Serving on port: ", viper.GetInt("port"))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntP("port", "p", 8080, "the web service port")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
}
