package config

import (
	"gospider/global/error"
	"gospider/global/variable"
	"github.com/spf13/viper"
	"log"
	"time"
)

// configYml configYml struct
type configYml struct {
	viper *viper.Viper
}

// CreateConfigYmlFactory create ymlfactory
func CreateConfigYmlFactory() *configYml {

	yamlConfig := viper.New()
	yamlConfig.AddConfigPath(variable.BasePath + "/config")
	// 需要读取的文件名
	yamlConfig.SetConfigName("config")
	//设置配置文件类型
	yamlConfig.SetConfigType("yaml")

	if err := yamlConfig.ReadInConfig(); err != nil {
		log.Fatal(error.ErrorsConfigInitFail + err.Error())
	}

	return &configYml{
		yamlConfig,
	}
}

// Get 一个原始值
func (c *configYml) Get(keyname string) interface{} {
	return c.viper.Get(keyname)
}

// GetString get string
func (c *configYml) GetString(keyname string) string {
	return c.viper.GetString(keyname)
}

// GetBool get bool
func (c *configYml) GetBool(keyname string) bool {
	return c.viper.GetBool(keyname)
}

// GetInt get int
func (c *configYml) GetInt(keyname string) int {
	return c.viper.GetInt(keyname)
}

// GetInt32 get int32
func (c *configYml) GetInt32(keyname string) int32 {
	return c.viper.GetInt32(keyname)
}

// GetInt64 get int64
func (c *configYml) GetInt64(keyname string) int64 {
	return c.viper.GetInt64(keyname)
}

// GetFloat64 get float64
func (c *configYml) GetFloat64(keyname string) float64 {
	return c.viper.GetFloat64(keyname)
}

// GetDuration get duration
func (c *configYml) GetDuration(keyname string) time.Duration {
	return c.viper.GetDuration(keyname)
}

// GetStringSlice get stringslice
func (c *configYml) GetStringSlice(keyname string) []string {
	return c.viper.GetStringSlice(keyname)
}