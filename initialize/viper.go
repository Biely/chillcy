package initialize

import (
	"fmt"

	"github.com/Biely/chillcy/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//定义配置文件名常量，读取配置文件
//@author: [Gaayean]

const (
	// ConfigEnv         = "RATE_CONFIG"
	ConfigDefaultFile = "config.yaml"
	// ConfigTestFile    = "config.test.yaml"
	// ConfigDebugFile   = "config.debug.yaml"
	// ConfigReleaseFile = "config.release.yaml"
)

func Viper() *viper.Viper {
	//可根据不同环境读取不同的配置文件
	//code..

	v := viper.New()
	v.SetConfigFile(ConfigDefaultFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.RATE_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.RATE_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
