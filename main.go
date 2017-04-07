package main

import (
    . "indie/core"
    "fmt"
    "indie/modules/stat"
    "os"
    "log"
)

func main() {
    log.Println(fmt.Sprintf("Start pid: %d", os.Getpid()))
    svr := NewIndie()

    // register Modules
    svr.RegModule([]IModule{
	// example module
	//*exam.NewExamModule(),

	// stat module
	*stat.NewStatModule(),

	// identify porn image module

    })

    svr.Serve()

    log.Println("[INFO] IMGEYE EXIT 0")
}
