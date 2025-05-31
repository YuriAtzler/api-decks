package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// instância única de *gorm.DB*
var db *gorm.DB

// Initialize abre a conexão uma vez e configura o pool
func Initialize() {
	if db != nil {
		return
	}

	dsn := "postgresql://postgres.icgzbdzjeouftniexhku:myNbu8-xyxsyn-votnor@aws-0-sa-east-1.pooler.supabase.com:6543/postgres"

	// aqui ativamos o simple protocol para evitar prepared-statements
	dialector := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})

	var err error
	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// extrai o *sql.DB* para configurar o pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// opcional: garante que está vivo
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}
}

// Get retorna a instância única de *gorm.DB*
// Chame Initialize() no seu main() antes de tudo
func Get() *gorm.DB {
	if db == nil {
		Initialize()
	}
	return db
}

// Close fecha o pool ao final da aplicação
func Close() error {
	if db == nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
