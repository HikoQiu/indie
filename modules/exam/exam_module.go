package exam

import (
    "indie/core"
    "fmt"
    "time"
    "reflect"
)

/**
 * 1. inherit from core.Module
 * 2. implement IModule interface
 */
type ExamModule struct {
    // inherit from core.Module
    core.Module
}

// implement IModule interface
func (t ExamModule) Config(conf *core.ModuleConf) {
    workerNum, _ := conf.Int("worker_num")
    fmt.Println("Config " + t.Name + "'s configuratons worker_num: ", workerNum)
}

// Module start to work
func (t ExamModule) Handle() {
    fmt.Println(reflect.TypeOf(t).String())
    fmt.Println(t.Name + " start to run...")
    time.Sleep(time.Second * 1)
}

// When receive TERM signal, Stop func will be invoked
func (t ExamModule) Stop() {
    fmt.Println(t.Name + " stop")
}

func NewExamModule() *ExamModule {
    m := new(ExamModule)

    // Module name, and [MODULE_EXAM] configuration section will be parse
    m.Name = "EXAM"
    return m
}
