package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string
	ConsumerKey string
	AuthServerUrl string
	TelegramBotUrl string `mapstructure:"bot_url"`
	DBPath string `mapstructure:"db_file"`

	Messages Messages

}

type Messages struct {
	Errors
	Responses
}

type Errors struct {
	Default string `mapstructure:"default"`
	InvalidURL string `mapstructure:"invalid_url"`
	Unauthorized string `mapstructure:"unauthorized"`
	UnableToSave string `mapstructure:"unable_to_save"`
}

type Responses struct {
	Start string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	SavedSuccessfully string `mapstructure:"saved_successfully"`
	UnknownCommand string `mapstructure:"unknown_command"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	os.Setenv("TOKEN", "7523392776:AAGtqeGkhbYNMKEpcXYAYoRtDMwhWbFZmu4")
	os.Setenv("CONSUMER_KEY", "112096-5338ab51ceb2b97df76480a")
	os.Setenv("AUTH_SERVER_URL", "http://localhost/")

	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	if err := viper.BindEnv("consumer_key"); err != nil {
		return err
	}

	if err := viper.BindEnv("auth_server_url"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("token")
	cfg.ConsumerKey = viper.GetString("consumer_key")
	cfg.AuthServerUrl = viper.GetString("auth_server_url")

	return nil
}