package mytime

import "time"

func GetEpochTime() int64 {
	return time.Now().Unix()
}

func GetUTCTimeFromEpoch(epoch int64) time.Time {
	return time.Unix(epoch, 0).UTC()
}
