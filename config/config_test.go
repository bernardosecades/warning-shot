package config

import (
	"testing"
)

func TestDefaultValuesWebServer(t *testing.T) {

	cfg := GetConfig()

	if cfg.WebServer.Port != "3000" {
		t.Error("Expected 3000, got", cfg.WebServer.Port)
	}
}

func TestDefaultValuesMySql(t *testing.T) {

	cfg := GetConfig()

	if cfg.MySql.Host != "localhost" {
		t.Error("Expected localhost, got", cfg.MySql.Host)
	}

	if cfg.MySql.Port != "3306" {
		t.Error("Expected 3306, got", cfg.MySql.Port)
	}

	if cfg.MySql.User != "root" {
		t.Error("Expected root, got", cfg.MySql.User)
	}

	if cfg.MySql.Password != "root" {
		t.Error("Expected root, got", cfg.MySql.Password)
	}

	if cfg.MySql.Database != "myDb" {
		t.Error("Expected myDb, got", cfg.MySql.Database)
	}
}
