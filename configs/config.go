package configs

var cfg *config

type config struct {
	DBDrive string
	DBHost  string
	DBName  string
	DBPort  string
	DBUser  string
	DBPass  string

	ServerPort string

	JWTSecret    string
	JWTExpiresIn int
}

// func Init(path string) (*config, error) {

// }
