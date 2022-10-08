package config

type Notifier_config struct {
	Config struct {
		Bot_Api    string `yaml:"telegram_bot_token"`
		Bot_Chanel int64  `yaml:"telegram_chanel"`
	} `yaml:"bot"`
}

type Redis_config struct {
	Config struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		Database    string `yaml:"database"`
		Path_Writer string `yaml:"path"`
	} `yaml:"redis"`
}

type Mongo_config struct {
	Config struct {
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mongo"`
}
