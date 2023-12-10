package Utils

import "github.com/go-ini/ini"

func GetStat(section string, key string) string {
	cfg, err := ini.Load("Servers/mysql/bin/my.ini")
	if err != nil {
		panic(err)
	}

	return cfg.Section(section).Key(key).String()
}
