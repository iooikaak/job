package util

import (
	"fmt"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(TimeLayout))
	return []byte(stamp), nil
}

func (t *JsonTime) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var err error
	i, err := time.ParseInLocation("2006-01-02 15:04:05", string(data), time.Local)
	*t = JsonTime(i)
	return err
}

func (t JsonTime) Empty() bool {
	return t == JsonTime{}
}
