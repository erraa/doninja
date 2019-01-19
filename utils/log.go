package utils

import (
	log "github.com/sirupsen/logrus"
)

func LogWithPrefix(prefix string) *log.Entry {
	return log.WithField("prefix", prefix)
}
