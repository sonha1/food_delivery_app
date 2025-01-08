package setting

type Config struct {
	Server ServerSetting `json:"server"`
	Auth   AuthSetting   `json:"auth"`
	Mysql  MySQLSetting  `json:"mysql"`
	Logger LoggerSetting `json:"logger"`
	Redis  RedisSetting  `json:"redis"`
	JWT    JWTSetting    `json:"jwt"`
}

type ServerSetting struct {
	Port int    `json:"port"`
	Mode string `json:"mode"`
}

type RedisSetting struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

type MySQLSetting struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Dbname          string `json:"dbname"`
	MaxIdleConns    int    `json:"maxIdleConns"`
	MaxOpenConns    int    `json:"maxOpenConns"`
	ConnMaxLifetime int    `json:"connMaxLifetime"`
}

type LoggerSetting struct {
	Log_level     string `json:"log_level"`
	File_log_name string `json:"file_log_name"`
	Max_backups   int    `json:"max_backups"`
	Max_age       int    `json:"max_age"`
	Max_size      int    `json:"max_size"`
	Compress      bool   `json:"compress"`
}

// JWT Settings
type JWTSetting struct {
	TOKEN_HOUR_LIFESPAN uint   `json:"TOKEN_HOUR_LIFESPAN"`
	API_SECRET_KEY      string `json:"API_SECRET_KEY"`
	JWT_EXPIRATION      string `json:"JWT_EXPIRATION"`
}

// Auth Setting
type AuthSetting struct {
	SaltRounds int `json:"salt_rounds"`
}
