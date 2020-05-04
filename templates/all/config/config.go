package config

type Config struct {
	Default  DefaultConfig `toml:"default"`
	Log      LogConfig     `toml:"log"`
	MasterDB PgConfig      `toml:"masterDB"`
	Redis    RedisConfig   `toml:"redis"`
}

type DefaultConfig struct {
	AppName      string
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

type LogConfig struct {
	Level string
}

type PgConfig struct {
	Host        string
	Port        int
	User        string
	Password    string
	DB          string
	MaxOpenConn int
	MaxIdleConn int
}

type RedisConfig struct {
	Addr        string
	Password    string
	DB          int
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}
