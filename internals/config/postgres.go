package config

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type PostgresClient struct {
	Read  *bun.DB
	Write *bun.DB
}

func NewPostgres() (*PostgresClient, func(), error) {
	readDns := os.Getenv("POSTGRES_READ_DNS")
	readDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(readDns)))
	read := bun.NewDB(readDb, pgdialect.New())
	writeDns := os.Getenv("POSTGRES_WRITE_DNS")
	writeDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(writeDns)))
	write := bun.NewDB(writeDb, pgdialect.New())

	if os.Getenv("GO_ENV") == "development" {
		write.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
		read.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return &PostgresClient{
			Read:  read,
			Write: write,
		}, func() {
			read.Close()
			write.Close()
		}, nil
}

// NewTestPostgres テスト接続用
func NewTestPostgres() (*PostgresClient, func(), error) {
	readDns := os.Getenv("POSTGRES_READ_DNS")
	readDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(readDns)))
	read := bun.NewDB(readDb, pgdialect.New())
	writeDns := os.Getenv("POSTGRES_WRITE_DNS")
	writeDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(writeDns)))
	write := bun.NewDB(writeDb, pgdialect.New())

	return &PostgresClient{
			Read:  read,
			Write: write,
		}, func() {
			read.Close()
			write.Close()
		}, nil
}
