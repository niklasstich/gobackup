/*
Copyright © 2021 niklasstich

*/
package cmd

import (
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gobackup",
	Short: "gobackup is a snapshotting backup utility",
	Long: `gobackup is a snapshotting backup utility similar to rnsapshot,
but without support for incremental backups meaning it works for filesystems that do not support hardlinks.`,
}

var timeformat = "Monday 2006/01/02 15:04:05 -0700"

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gobackup.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("gobackup.yaml")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath("$HOME/.config/")
	viper.AddConfigPath("$HOME/.config/gobackup")
	viper.AddConfigPath("$HOME/gobackup/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
}
