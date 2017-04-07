package core

import "github.com/larspensjo/config"

const (
    // Prefix in config file's section
    MODULE_CONF_PREFIX = "MODULE_"
)

// Define module conf
// How to use in module?
// e.g: val, _ :=  moduleConf.Int("field_name")
// NOTICE: before using it, please config the [MODULE_{MODULE_NAME}] configuration section first.
// e.g:
// [MODULE_EXAM]
// field_name1 = 1
type ModuleConf struct {
    config     *config.Config
    ModuleName string
}

func NewModuleConf(name string, config *config.Config) *ModuleConf {
    c := new(ModuleConf)
    c.config = config
    c.ModuleName = name
    return c
}

func (mc *ModuleConf) Int(option string) (int, error) {
    return mc.config.Int(MODULE_CONF_PREFIX + mc.ModuleName, option)
}

func (mc *ModuleConf) String(option string) (string, error) {
    return mc.config.String(MODULE_CONF_PREFIX + mc.ModuleName, option)
}

func (mc *ModuleConf) Float(option string) (float64, error) {
    return mc.config.Float(MODULE_CONF_PREFIX + mc.ModuleName, option)
}

func (mc *ModuleConf) Bool(option string) (bool, error) {
    return mc.config.Bool(MODULE_CONF_PREFIX + mc.ModuleName, option)
}
