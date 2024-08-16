package cmd

import (
	"net/http"

	"github.com/Adosh74/geek-blog/pkg/config"
	"github.com/Adosh74/geek-blog/pkg/routing"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  "Application will be served on host and port defined in config.yaml file",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {

	config.Set()

	router := routing.GetRouter()

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "OK",
			"app name": viper.Get("App.Name"),
		})
	})
	routing.Serve()
}
