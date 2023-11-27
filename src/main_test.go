package main

import (
	"family-team/src/config"
	"testing"
)

func TestParseConfig(t *testing.T) {
	c := config.ParseConfig()
	if c.Port != "80" {
		t.Errorf("Expected port to be 80, got %s", c.Port)
	}
	if c.File != "./test.zip" {
		t.Errorf("Expected file to be ./test.zip, got %s", c.File)
	}
	if c.Ext != ".c" {
		t.Errorf("Expected ext to be .c, got %s", c.Ext)
	}
}
