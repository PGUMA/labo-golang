package config

import "os"

func Value(key string) string {
	v, e := os.LookupEnv(key)
	if !e {
		return "not set"
	} else {
		return v
	}
}
