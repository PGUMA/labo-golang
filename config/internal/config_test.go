package config_test

import (
	"testing"

	"github.com/PGUMA/labo-golang/config"
)

func TestValue(t *testing.T) {
	key := "key"
	t.Setenv(key, "test_env")

	v := config.Value(key)
	if v != "test_env" {
		t.Errorf("want %s, but %s", "test_env", v)
	}
}
