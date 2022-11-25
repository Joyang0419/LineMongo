package configs

import (
	"LineMongo/internals/tools"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type Config struct {
	MongoDBSetting *tools.MongoDBSetting
	LineBotSetting *tools.LineBotSetting
}

func LoadConfig(path string, configName string, configType string) (config Config, err error) {
	if lastPath := filepath.Base(path); lastPath == "integrations" {
		path = strings.ReplaceAll(path, "test/integrations", "")
	}
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	var (
		dbConf      *tools.MongoDBSetting
		lineBotConf *tools.LineBotSetting
	)
	err = viper.Unmarshal(&dbConf)
	err = viper.Unmarshal(&lineBotConf)
	config.MongoDBSetting = dbConf
	config.LineBotSetting = lineBotConf
	return config, nil
}
