package src

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config *BackupConfig

type BackupConfig struct {
	BackupDestination         string
	BackupSources             []string
	BackupSourcesNonRecursive []string
	RetainStrategy            struct {
		Daily   int
		Weekly  int
		Monthly int
	}
}

func GetConfig() (c *BackupConfig) {
	if Config == nil {
		getConfigInternal()
	}
	return Config
}

func getConfigInternal() {
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("failed to unmarshal config into struct: %v", err)
	}
}
