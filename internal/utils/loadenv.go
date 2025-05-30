package utils

import (
	"os"
)

type EnvVar struct {
	Key          string
	DefaultValue string
}

func LoadEnv(vars ...EnvVar) error {
	for _, v := range vars {
		value, exists := os.LookupEnv(v.Key)
		if !exists || value == "" {
			if v.DefaultValue != "" {
				os.Setenv(v.Key, v.DefaultValue)
			} else {
				return &os.PathError{Op: "loadenv", Path: v.Key, Err: os.ErrNotExist}
			}
		}
	}
	return nil
}
