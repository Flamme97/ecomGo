package config

import ("os"
"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUSer     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()


func initConfig() Config{
	godotenv.Load()
	return Config{
		PublicHost: getEnv("Public_host", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBUSer: getEnv("DB_USer", "root"),
		DBPassword: getEnv("DB_PASSWORD", "myPassword"),
		DBAddress: getEnv("DB_ADDRESS", "5432"),
		DBName: getEnv("DB_NAME", "ecom"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}