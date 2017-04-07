package core

import (
    "github.com/larspensjo/config"
    "log"
)

const (
    APP = "indie"
)

// App Configuration
type Conf struct {
    // configuration reader
    C       *config.Config

    // Configuration fields relative to config.ini
    Version string
}

// get config obj
func GetConf(path string) (*Conf, error) {
    c, err := config.ReadDefault(path)
    if err != nil {
	log.Fatalln("[ERROR] ERROR MSG: ", err.Error())
	return nil, err
    }
    conf := parseConf(c)
    conf.C = c
    return conf, nil
}

// parse app config
func parseConf(c *config.Config) *Conf {
    conf := new(Conf)
    conf.Version, _ = c.String(APP, "version")
    return conf
}

