package src

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

func RotationMonthly() {
	config := GetConfig()
	rotationGeneric(config.BackupDestination, "monthly", config.RetainStrategy.Monthly)
}

func RotationWeekly() {
	config := GetConfig()
	rotationGeneric(config.BackupDestination, "weekly", config.RetainStrategy.Weekly)
}

func RotationDaily() {
	config := GetConfig()
	rotationGeneric(config.BackupDestination, "daily", config.RetainStrategy.Daily)
}

//rotates all [interval].0 ... [interval].maxRetention-1 up by one and frees [interval].0 in doing so
func rotationGeneric(destRoot, intervalName string, maxRetention int) {
	_, err := os.Stat(destRoot)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("directory %v specified for BackupDestination does not exist: %v", destRoot, err)
		}
		log.Fatalf("failed to stat directory %v specified for BackupDestination: %v", destRoot, err)
	}
	pathTemplate := path.Join(destRoot, intervalName+".%d")
	for i := maxRetention - 1; i > 0; i-- {
		newPath := fmt.Sprintf(pathTemplate, i)
		oldPath := fmt.Sprintf(pathTemplate, i-1)

		//remove the directory we are trying to move to
		err := os.RemoveAll(newPath)
		if err != nil {
			log.Fatalf("failed to remove newPath %v: %v", newPath, err)
		}

		//if the directory we are trying to move back one doesn't exist, just continue
		if _, err := os.Stat(oldPath); os.IsNotExist(err) {
			continue
		}

		//move 0 to 1, 1 to 2 and so on
		err = os.Rename(oldPath, newPath)
		if err != nil {
			log.Fatalf("failed to move %v to %v: %v", oldPath, newPath, err)
		}
	}
}
