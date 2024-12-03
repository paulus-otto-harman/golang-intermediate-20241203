package util

import (
	"strconv"
)

func Uint(param string) (uint, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
