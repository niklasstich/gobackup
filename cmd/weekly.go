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

// weeklyCmd represents the weekly command
var weeklyCmd = &cobra.Command{
	Use:   "weekly",
	Short: "Perform weekly backup according to current configuration",
	Long: `Performs the weekly backup according to your current configuration, meaning it will move all existing backups
back by one week (weekly.0 becomes weekly.1, weekly.1 becomes weekly.2 and so on) and creates a new backup in weekly.0

The backup weekly.[RetainStrategy.Weekly-1] will be deleted by this action.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Starting weekly backup at %v", time.Now().Format(timeformat))
		src.RotationWeekly()
		src.BackupWeekly()
		log.Infof("Finished weekly backup at %v", time.Now().Format(timeformat))
	},
}

func init() {
	rootCmd.AddCommand(weeklyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weeklyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weeklyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
