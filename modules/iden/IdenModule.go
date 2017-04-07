package iden

import "indie/core"

type IdenModule struct {
    core.Module
}

func (t IdenModule) Config(conf *core.ModuleConf) {

}

func (t IdenModule) Handle() {

}

func (t IdenModule) Stop() {

}

func NewIdenModule() *IdenModule {
    i := new(IdenModule)
    i.Name = "IDEN"
    return i
}
