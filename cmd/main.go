package main

import (
	"database/sql"
	"log"
	"migrations-service/internal/core/db/migrations"
	"migrations-service/internal/core/settings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func main() {
	settings.LoadDotEnv()

	settings.LoadEnvs()

	// Configurar pool de conexiones
	db, err := sql.Open("postgres", settings.Settings.DATABASE_URL)
	if err != nil {
		log.Fatalf("Error al abrir la conexión a la base de datos: %v", err)
	}

	// Validar conexión al inicio
	if err := db.Ping(); err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	log.Println("Conexión a la base de datos establecida")

	// Ejecutar el servidor
	defer db.Close()

	if err := runMigrations(db); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
}

func runMigrations(db *sql.DB) error {
	sourceDriver, err := iofs.New(migrations.Files, ".")
	if err != nil {
		return err
	}

	// Crear el adaptador para la base de datos PostgreSQL
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	// Crear instancia de migrate
	m, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", dbDriver)
	if err != nil {
		return err
	}

	// Ejecutar migraciones
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migraciones aplicadas exitosamente")
	return nil
}
