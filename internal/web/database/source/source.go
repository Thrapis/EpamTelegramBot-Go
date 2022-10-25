package source

import (
	cfg "baa-telebot/internal/config"
	"fmt"
	"sync"
)

var source string
var load sync.Once

func GetSource() string {
	load.Do(func() {
		config := cfg.GetConfig()
		source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			config.Database.User, config.Database.Password, config.Database.Host,
			config.Database.Port, config.Database.Instance)
	})
	return source
}
