package database

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"pipeliner/config"
)

type Db struct {
	conn *pgxpool.Pool
}

func GetDb() (Db, error) {
	conn, err := pgxpool.Connect(context.Background(),
		"postgres://"+config.Config.Database.User+":"+
			config.Config.Database.Password+"@"+config.Config.Database.Host+":"+
			config.Config.Database.Port+"/"+config.Config.Database.DbName)

	db := Db{conn: conn}
	return db, err
}

func (db *Db) Query(query string) error {
	rows, err := db.conn.Query(context.Background(), query)
	rows.Close()
	if err != nil {
		return err
	}

	return nil
}

// QueryStruct structPointer is pointer Struct array. Example *[]User
func (db *Db) QueryStruct(query string, structPointer interface{}) error {
	err := pgxscan.Select(context.Background(), db.conn, structPointer, query)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) QueryRows(query string) (pgx.Rows, error) {
	rows, err := db.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (db *Db) QueryRow(query string) pgx.Row {
	row := db.conn.QueryRow(context.Background(), query)
	return row
}

func (db *Db) Close() {
	db.conn.Close()
}
