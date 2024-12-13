package settings

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DATABASE_URL string `required:"true"`
}

var Settings Config
var EnvDir = ".envs"

func LoadDotEnv() {

	err := godotenv.Load(fmt.Sprintf("%s/.env.base", EnvDir))
	if err != nil {
		log.Printf("No %s file found, using system environment variables", fmt.Sprintf("%s/.env.base", EnvDir))
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		log.Println("ENVIRONMENT is not set")
	}

	// Mapear el archivo .env correspondiente al entorno
	envFiles := map[string]string{
		"":            fmt.Sprintf("%s/.env", EnvDir),
		"local":       fmt.Sprintf("%s/.env.local", EnvDir),
		"development": fmt.Sprintf("%s/.env.dev", EnvDir),
		"production":  fmt.Sprintf("%s/.env.prod", EnvDir),
		"staging":     fmt.Sprintf("%s/.env.staging", EnvDir),
	}

	// Obtener el archivo de entorno correspondiente
	envFile, exists := envFiles[environment]
	if !exists {
		log.Fatalf("Environment '%s' is not supported. Must be one of: local, development, production, staging", environment)
	}

	// Cargar las variables desde el archivo correspondiente
	err = godotenv.Load(envFile)
	if err != nil {
		log.Printf("No %s file found, using system environment variables", envFile)
	} else {
		log.Printf("Loaded environment variables from %s", envFile)
	}
}

func LoadEnvs() {
	// Procesar las variables de entorno en la estructura Settings
	err := envconfig.Process("", &Settings)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
}
