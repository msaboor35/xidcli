package xid

import (
	"time"

	"github.com/rs/xid"
)

func Time(xidStr string) (*time.Time, error) {
	x, err := xid.FromString(xidStr)
	if err != nil {
		return nil, err
	}

	time := x.Time()
	return &time, nil
}

func Machine(xidStr string) ([]byte, error) {
	x, err := xid.FromString(xidStr)
	if err != nil {
		return []byte{}, err
	}

	return x.Machine(), nil
}

func Pid(xidStr string) (uint16, error) {
	x, err := xid.FromString(xidStr)
	if err != nil {
		return 0, err
	}

	return x.Pid(), nil
}

func Counter(xidStr string) (int32, error) {
	x, err := xid.FromString(xidStr)
	if err != nil {
		return 0, err
	}

	return x.Counter(), nil
}

func UnixTime(xidStr string) (int64, error) {
	x, err := xid.FromString(xidStr)
	if err != nil {
		return 0, err
	}

	return x.Time().Unix(), nil
}
