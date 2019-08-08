package config

import (
    "github.com/suratu-io/curl-telegram-bot/app/logger"
    "os"
    "strconv"
)

const (
    // BotApiTokenEnv is env var key for BotApiToken.
    BotApiTokenEnv   = "BOT_API_TOKEN"
    // UpdateTimeoutEnv is env var key for UpdateTimeout.
    UpdateTimeoutEnv = "UPDATE_TIMEOUT"
)

var (
    log = logger.Factory.NewLogger("config")

    // BotApiToken is used to connect to bot api.
    BotApiToken   = os.Getenv(BotApiTokenEnv)
    // UpdateTimeout is a duration in seconds for bot api update chan timout.
    UpdateTimeout = 60
)

func init() {
    ut, err := strconv.ParseInt(os.Getenv(UpdateTimeoutEnv), 10, 32)
    if err != nil {
        log.Warn.Println("update timeout has not been set properly, using the default")
    } else {
        UpdateTimeout = int(ut)
    }
}
