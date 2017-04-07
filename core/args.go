package core

import (
    "flag"
    "os"
    "strings"
    "errors"
    "log"
)

// Cli parameters
type Args struct {
    ConfigPath string
}

// Get parameter obj
func GetArgs() *Args {
    args := new(Args)
    h := flag.Bool("h", false, "Get help informations")
    c := flag.String("c", "./conf/config.ini", "Config file path")
    flag.Parse()
    if *h {
	flag.Usage()
	os.Exit(0)
    }

    args.ConfigPath = *c
    err := checkArgs(args)
    if err != nil {
	log.Fatalln("[ERROR] INVALID ARGS: " + err.Error())
    }

    return args
}

func checkArgs(args *Args) error {
    // Check whether config file exists
    if strings.Index(args.ConfigPath, "./") != 0 && strings.Index(args.ConfigPath, "/") != 0 {
	args.ConfigPath = "./" + args.ConfigPath
    }
    _, err := os.Stat(args.ConfigPath)
    if err != nil {
	return errors.New("CONFIG FILE NOT FOUND")
    }

    return nil
}

