/*
Copyright © 2021 niklasstich

*/
package cmd

import (
	"github.com/niklasstich/gobackup/src"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

// dailyCmd represents the daily command
var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Perform daily backup according to current configuration",
	Long: `Performs the daily backup according to your current configuration, meaning it will move all existing backups
back by one day (daily.0 becomes daily.1, daily.1 becomes daily.2 and so on) and creates a new backup in daily.0

The backup daily.[RetainStrategy.Daily-1] will be deleted by this action.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Starting daily backup at %v", time.Now().Format(timeformat))
		src.RotationDaily()
		src.BackupDaily()
		log.Infof("Finished daily backup at %v", time.Now().Format(timeformat))
	},
}

func init() {
	rootCmd.AddCommand(dailyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dailyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dailyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
