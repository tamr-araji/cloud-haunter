package utils

import (
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/hortonworks/cloud-cost-reducer/types"
)

func IsAnyMatch(haystack types.Tags, needles ...string) bool {
	for _, k := range needles {
		if _, ok := haystack[k]; ok {
			return true
		}
	}
	return false
}

func ConvertTimeRFC3339(stringTime string) (time.Time, error) {
	return time.Parse(time.RFC3339, stringTime)
}

// ConvertTimeUnix parses a unix timestamp (seconds since epoch start) from string to time.Time
func ConvertTimeUnix(unixTimestamp string) time.Time {
	timestamp, err := strconv.ParseInt(string(unixTimestamp), 10, 64)
	if err != nil {
		log.Warnf("[util.ConvertTimeUnix] cannot convert time: %s, err: %s", unixTimestamp, err)
		return time.Unix(0, 0)
	}
	return time.Unix(timestamp, 0)
}

type ConvertTagsFuncSignature func(map[string]*string) types.Tags

func ConvertTags(tagMap map[string]*string) types.Tags {
	tags := make(types.Tags, 0)
	for k, v := range tagMap {
		tags[k] = *v
	}
	return tags
}