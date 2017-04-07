package core

import (
    "sync"
    "os"
    "log"
    "reflect"
    "os/signal"
    "syscall"
)

type Indie struct {
    // App Configuration
    conf       *Conf

    // modules
    modules    []IModule

    // modules' WaitGroup, for graceful kill
    // Two places where moduleWgs[{MODULE_NAME}].Done() is invoke
    // 1. When m.Handle() finishes normally
    // 2. Progress receives a SIGTERM signal
    moduleWgs  map[string]*sync.WaitGroup

    // for accepting system signal
    signalChan chan os.Signal
}

func NewIndie() (*Indie) {
    i := new(Indie)
    i.conf, _ = GetConf(GetArgs().ConfigPath)
    i.signalChan = make(chan os.Signal)
    i.moduleWgs = make(map[string]*sync.WaitGroup)
    return i
}

// register modules
func (t *Indie) RegModule(modules []IModule) {
    for _, m := range modules {
	if flag, err := inheritModule(m); !flag {
	    log.Fatalln(err.Error())
	}

	t.modules = append(t.modules, m)
	log.Println("[INFO] MODULE - \"" + reflect.TypeOf(m).String() + "\" REGISTERED ... OK")
    }
}

// init registered modules
func (t *Indie) initModules() {
    for _, m := range t.modules {
	t.runModule(m)
    }
}

// run specify module
func (t *Indie) runModule(m IModule) {
    // reflect to get module info
    im := reflect.ValueOf(m).FieldByName("Module").Interface()
    if im == nil {
	panic("nil module")
    }

    // set configuration for module
    tm := im.(Module)
    m.Config(NewModuleConf(tm.Name, t.conf.C))

    // set module WaitGroup, for graceful kill
    t.moduleWgs[tm.Name] = &sync.WaitGroup{}
    t.moduleWgs[tm.Name].Add(1)

    // run module Handle() to start module
    go func() {
	defer t.moduleWgs[tm.Name].Done()
	m.Handle()
    }()
}

// for graceful kill. Here handle SIGTERM signal to do sth
// e.g: kill -TERM $pid
func (t *Indie) handleSignal() {
    signal.Notify(t.signalChan, syscall.SIGTERM)

    for {
	switch <-t.signalChan {
	case syscall.SIGTERM:
	    log.Println("[NOTICE] RECEIVE signal: SIGTERM. ")

	    // Notify Stop
	    t.stopModules()

	    // Wait Modules Stop
	    t.waitModules()

	    log.Println("[NOTICE] ALL MODULES ARE STOPPED, GOING TO EXIT.")

	    os.Exit(0)
	}
    }
}

// invoke module Stop() to stop accept new Jobs
func (t *Indie) stopModules() {
    for _, m := range t.modules {
	m.Stop()
	im := reflect.ValueOf(m).FieldByName("Module").Interface()
	t.moduleWgs[im.(Module).Name].Done()
    }
}

// wait module goroutines end
func (t *Indie) waitModules() {
    for _, m := range t.modules {
	im := reflect.ValueOf(m).FieldByName("Module").Interface()
	t.moduleWgs[im.(Module).Name].Wait()
    }
}

// wait modules to exit
func (t *Indie) Wait() {
    t.waitModules()
}

// Start to Serve
func (t *Indie) Serve() {

    // listen to os signal, for graceful kill
    go t.handleSignal()

    // init modules
    t.initModules()

    // wait till all goroutines exit
    t.Wait()

    log.Println("[INFO] PROC EXIT - ", os.Getpid())
}
