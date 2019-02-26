package cadb

import (
	"database/sql"
	"fmt"

	caErr "github.com/roosr/gotest/ca/error"
	"github.com/roosr/gotest/ca/model"
)

type CADao struct {
	db *sql.DB
}

func New() *CADao {
	dao := &CADao{
		db: dbConn(),
	}

	return dao
}

func (dao *CADao) GetName(id string) (*model.User, error) {

	rows, err := dao.db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return nil, caErr.New(caErr.DBError)
	}

	if !rows.Next() {

		err = rows.Err()
		if err != nil {
			fmt.Printf("%s", err.Error())
			return nil, caErr.New(caErr.DBError)
		}

		return nil, nil
	}

	var userId, username, email string
	err = rows.Scan(&userId, &username, &email)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return nil, caErr.New(caErr.DBError)
	}

	return &model.User{
		Id:       userId,
		Username: username,
		Email:    email,
	}, nil
}
