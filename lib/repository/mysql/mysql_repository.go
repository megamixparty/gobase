package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/megamixparty/gobase/lib/repository"
)

// MysqlRepository struct contains Database to access tables belong to Mysql
type MysqlRepository struct {
	db *sqlx.DB
}

// NewMysqlRepository returns database implementation of IUserRepository
func NewMysqlRepository(db *sqlx.DB) *MysqlRepository {
	return &MysqlRepository{db}
}

// GetConnection returns appropiate DB connection object
func (mr *MysqlRepository) GetConnection(ctx context.Context, transaction bool) (sqlx.ExtContext, error) {
	if transaction {
		tx, err := mr.db.BeginTxx(ctx, nil)
		if err != nil {
			return nil, err
		}
		return tx, nil
	}
	return mr.db, nil
}

// Commit commits SQL queries into database
func (mr *MysqlRepository) Commit(db sqlx.ExtContext) error {
	err := db.(*sqlx.Tx).Commit()
	if err != nil {
		return err
	}

	return nil
}

// Rollback rollbacks SQL queries from database
func (mr *MysqlRepository) Rollback(db sqlx.ExtContext) error {
	err := db.(*sqlx.Tx).Rollback()
	if err != nil {
		return err
	}

	return nil
}

// NewUserRepository returns UserMysql object
func (mr *MysqlRepository) NewUserRepository(db sqlx.ExtContext) repository.IUserRepository {
	return NewUserMysql(db)
}
