package setting

type Config struct {
	Mysql  MySQLSetting `mapstructure:"mysql"`
	Logger LogSetting   `mapstructure:"log"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    string `mapstructure:"maxIdleConns"`
	MaxOpenConns    string `mapstructure:"maxOpenConns"`
	ConnMaxLifetime string `mapstructure:"connMaxLifetime"`
}

type LogSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackup   int    `mapstructure:"max_backup"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}
