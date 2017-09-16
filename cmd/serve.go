package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/judy2k/mudio/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"server", "run"},
	Short:   "Run the mudio web service",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt("port")
		fmt.Println("Serving on port: ", port)
		server.Run(port)
	},
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	RootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntP("port", "p", 8080, "the web service port")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
}
