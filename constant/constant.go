package constant

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	// ProxyPort 对外接口
	ProxyPort = "ProxyPort"
)

// Viper 全局配置
var Viper *viper.Viper = NewViper()

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("inig config err:%v", err.Error()))
	}
}

// NewViper 获取全局Viper
func NewViper() *viper.Viper {
	return viper.GetViper()
}
