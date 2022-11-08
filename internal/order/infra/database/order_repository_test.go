package database

import (
	"database/sql"
	"github.com/stretch/testify/suite"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	
}