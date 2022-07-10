package config

import "time"

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Database   DatabaseConfig   `mapstructure:"database"`
	Monitoring MonitoringConfig `mapstructure:"monitoring"`
	Redis      RedisConfig      `mapstructure:"redis"`
}

type ServerConfig struct {
	Environment        string
	AllowedOrigins     []string      `mapstructure:"allowed_origins"`
	AllowedHeaders     []string      `mapstructure:"allowed_headers"`
	AllowedMethod      []string      `mapstructure:"allowed_method"`
	Host               string        `mapstructure:"host"`
	UploadPath         string        `mapstructure:"upload_path"`
	MaxAgeHour         time.Duration `mapstructure:"max_age_hour"`
	SwaggerEndpoint    string        `mapstructure:"swagger_endpoint"`
	StaticFileEndpoint string        `mapstructure:"static_file_endpoint"`
	Port               int           `mapstructure:"port"`
	UploadLimitMB      int           `mapstructure:"upload_limit_mb"`
	Debug              bool          `mapstructure:"debug"`
	CorsEnable         bool          `mapstructure:"cors_enable"`
	AllowCredentials   bool          `mapstructure:"allow_credentials"`
}

type DatabaseConfig struct {
	MaxLifeTimeSec  time.Duration `mapstructure:"max_life_time_sec"`
	MasterDBAddress string        `mapstructure:"master_address"`
	MasterDBName    string        `mapstructure:"master_db_name"`
	SlaveDBAddress  string        `mapstructure:"slave_address"`
	SlaveDBName     string        `mapstructure:"slave_db_name"`
	MasterDBUname   string        `mapstructure:"master_db_uname"`
	MasterDBPass    string        `mapstructure:"master_db_pass"`
	SlaveDBUname    string        `mapstructure:"slave_db_uname"`
	SlaveDBPass     string        `mapstructure:"slave_db_pass"`
	MaxPoolOpen     int           `mapstructure:"max_pool_open"`
	MaxPoolIdle     int           `mapstructure:"max_pool_idle"`
	Debug           bool          `mapstructure:"debug"`
}

type MonitoringConfig struct {
	LogPath        string `mapstructure:"log_path"`
	LogName        string `mapstructure:"log_name"`
	TracingUrl     string `mapstructure:"tracing_url"`
	TracingAppName string `mapstructure:"tracing_app_name"`
}

type RedisConfig struct {
	Host       string        `mapstructure:"host"`
	Pass       string        `mapstructure:"pass"`
	TimeoutSec time.Duration `mapstructure:"timeout_sec"`
	PoolSize   int           `mapstructure:"pool_size"`
	MaxActive  int           `mapstructure:"max_active"`
	MaxIdle    int           `mapstructure:"max_idle"`
}
