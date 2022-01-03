package src

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func BackupMonthly() {
	config := GetConfig()
	backupGeneric(config.BackupDestination, "monthly", config.BackupSources, config.BackupSourcesNonRecursive)
}

func BackupWeekly() {
	config := GetConfig()
	backupGeneric(config.BackupDestination, "weekly", config.BackupSources, config.BackupSourcesNonRecursive)
}

func BackupDaily() {
	config := GetConfig()
	backupGeneric(config.BackupDestination, "daily", config.BackupSources, config.BackupSourcesNonRecursive)
}

func backupGeneric(destRoot, intervalName string, sources []string, nonRecursiveSources []string) {
	if !Exists(destRoot) {
		log.Warning("directory %v specified for BackupDestination does not exist - creating it")
		err := os.Mkdir(destRoot, os.ModeDir|0755)
		if err != nil {
			log.Fatalf("failed to create directory at %v: %v", destRoot, err)
		}
	}

	newPath := path.Join(destRoot, intervalName+".0")
	err := os.Mkdir(newPath, os.ModeDir|0755)
	if err != nil {
		log.Fatalf("failed to create directory at %v: %v", newPath, err)
	}

	for _, source := range sources {
		if err = CopyRecursive(source, newPath); err != nil {
			log.Errorf("Failed to backup %v to %v: %v", source, destRoot, err)
		}
	}

	for _, source := range nonRecursiveSources {
		CopyNonRecursive(source, destRoot)
	}

	//touch timestamp file
	filename := time.Now().Format("2006-01-02-15:04:05")
	filepath := path.Join(newPath, filename)
	f, err := os.Create(filepath)
	if err != nil {
		log.Errorf("Failed to create timestamp file %v in %v: %v", filename, filepath, err)
		return
	}
	defer f.Close()
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
