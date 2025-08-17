package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	DBUrl          string
	APIPort        string
	JWTSecret      string
	JWTExpire      time.Duration
	GRPCTimeout    time.Duration
	AuthService    string
	UserService    string
	GroupService   string
	MemberService  string
	ExpenseService string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	jwtExpire, err := time.ParseDuration(os.Getenv("JWT_EXPIRE"))
	if err != nil {
		log.Fatalf("Invalid JWT_EXPIRE: %v", err)
	}

	grpcTimeout, err := time.ParseDuration(os.Getenv("GRPC_TIMEOUT"))
	if err != nil {
		log.Fatalf("Invalid GRPC_TIMEOUT: %v", err)
	}

	return &Config{
		DBUrl:          os.Getenv("DB_URL"),
		APIPort:        os.Getenv("API_PORT"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		JWTExpire:      jwtExpire,
		GRPCTimeout:    grpcTimeout,
		AuthService:    os.Getenv("AUTH_SERVICE_ADDR"),
		UserService:    os.Getenv("USER_SERVICE_ADDR"),
		GroupService:   os.Getenv("GROUP_SERVICE_ADDR"),
		MemberService:  os.Getenv("MEMBER_SERVICE_ADDR"),
		ExpenseService: os.Getenv("EXPENSE_SERVICE_ADDR"),
	}
}
