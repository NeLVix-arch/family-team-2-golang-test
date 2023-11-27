package config

import "flag"

// Config is a structure that holds the parameters of the command line
type Config struct {
	Port string // port for HTTP server
	File string // path to ZIP file
	Ext  string // extension of files in ZIP archive
}

// ParseConfig returns an instance of Config from the flags
func ParseConfig() *Config {
	c := &Config{}
	flag.StringVar(&c.Port, "port", "80", "port for HTTP server")
	flag.StringVar(&c.File, "file", "./test.zip", "path to ZIP file")
	flag.StringVar(&c.Ext, "ext", ".c", "extension of files in ZIP archive")
	flag.Parse()
	return c
}
