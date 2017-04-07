package core

type IModule interface {
    // Get configuration section [MODULE_{MODULE_NAME}], and pass through Config() func
    Config(conf *ModuleConf)

    // Module start to work
    Handle()

    // When receive TERM signal, Stop() func will be invoked
    Stop()
}
