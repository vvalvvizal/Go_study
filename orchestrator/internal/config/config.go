package config

type Config struct {
	Port int
	DBHost string
	DBPort int
	DBUser string
	DBPassword string
	DBName string
}

/*
func LoadConfig() (Config, error){  //struct 복사본 값 복사, 설정 없으면 error로
	
	return config, nill
}
*/
 
func LoadConfig() *Config {
	return &Config{
		Port:       8080,
		DBHost:     "localhost",
		DBPort:     3306,
		DBUser:     "root",
		DBPassword: "password",
		DBName:     "orchestrator",
	}
}

func Load() (*Config, error) {
	return LoadConfig(), nil
}