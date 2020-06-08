package config

import (
	"gospider/global/error"
	"gospider/global/variable"
	"log"
	"time"

	"github.com/spf13/viper"
)

// SettingYml SettingYml struct
type SettingYml struct {
	viper *viper.Viper
}

// CreateSettingYmlFactory create ymlfactory
func CreateSettingYmlFactory() *SettingYml {
	yamlConfig := viper.New()
	// config file basepath
	yamlConfig.AddConfigPath(variable.BasePath + "/config")
	// filename of config
	yamlConfig.SetConfigName("setting")
	// filetype of config
	yamlConfig.SetConfigType("yaml")

	if err := yamlConfig.ReadInConfig(); err != nil {
		log.Fatal(error.ErrorsConfigInitFail + err.Error())
	}
	return &SettingYml{
		yamlConfig,
	}
}

// Get get keyname
func (c *SettingYml) Get(keyname string) interface{} {
	return c.viper.Get(keyname)
}

// GetString get string
func (c *SettingYml) GetString(keyname string) string {
	return c.viper.GetString(keyname)
}

// GetBool get bool
func (c *SettingYml) GetBool(keyname string) bool {
	return c.viper.GetBool(keyname)
}

// GetInt get int
func (c *SettingYml) GetInt(keyname string) int {
	return c.viper.GetInt(keyname)
}

// GetInt32 get int32
func (c *SettingYml) GetInt32(keyname string) int32 {
	return c.viper.GetInt32(keyname)
}

// GetInt64 get int64
func (c *SettingYml) GetInt64(keyname string) int64 {
	return c.viper.GetInt64(keyname)
}

// GetFloat64 get float64
func (c *SettingYml) GetFloat64(keyname string) float64 {
	return c.viper.GetFloat64(keyname)
}

// GetDuration get duration
func (c *SettingYml) GetDuration(keyname string) time.Duration {
	return c.viper.GetDuration(keyname)
}

// GetStringSlice get stringslice
func (c *SettingYml) GetStringSlice(keyname string) []string {
	return c.viper.GetStringSlice(keyname)
}
