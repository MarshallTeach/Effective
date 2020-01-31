package cache_service

import (
	"Effective/pkg/e"
	"strconv"
	"strings"
)

type Info struct {
	ID int
}

func (i *Info) GetInfoKey() string {
	keys := []string{
		e.CACHE_BASIC_INFO,
		"LIST",
	}

	if i.ID > 0 {
		keys = append(keys, strconv.Itoa(i.ID))
	}

	return strings.Join(keys, "_")
}