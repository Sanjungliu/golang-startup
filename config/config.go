package config

const (
	Environment        = "ENVIRONMENT"
	DBConnectionString = "DB_CONNECTION_STRING"
	JWTSecretKey       = "JWT_SECRET_KEY"
	MidtransServerKey  = "MIDTRANS_SERVER_KEY"
	MidtransClientKey  = "MIDTRANS_CLIENT_KEY"
)

type Config struct{}

func Init() *Config {
	return &Config{}
}

func (c *Config) Environment() string {
	return getStringOrDefault(Environment, "development")
}

func (c *Config) DBConnectionString() string {
	return getStringOrDefault(DBConnectionString, "")
}

func (c *Config) JWTSecretKey() string {
	return getStringOrDefault(JWTSecretKey, "")
}

func (c *Config) MidtransServerKey() string {
	return getStringOrDefault(MidtransServerKey, "")
}

func (c *Config) MidtransClientKey() string {
	return getStringOrDefault(MidtransClientKey, "")
}
