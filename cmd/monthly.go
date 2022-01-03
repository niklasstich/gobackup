/*
Copyright © 2021 niklasstich

*/
package cmd

import (
	"github.com/niklasstich/gobackup/src"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/spf13/cobra"
)

// monthlyCmd represents the monthly command
var monthlyCmd = &cobra.Command{
	Use:   "monthly",
	Short: "Perform monthly backup according to current configuration",
	Long: `Performs the monthly backup according to your current configuration, meaning it will move all existing backups
back by one month (monthly.0 becomes monthly.1, monthly.1 becomes monthly.2 and so on) and creates a new backup in monthly.0

The backup monthly.[RetainStrategy.Monthly-1] will be deleted by this action.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Starting monthly backup at %v", time.Now().Format(timeformat))
		src.RotationMonthly()
		src.BackupMonthly()
		log.Infof("Finished monthly backup at %v", time.Now().Format(timeformat))
	},
}

func init() {
	rootCmd.AddCommand(monthlyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monthlyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monthlyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
