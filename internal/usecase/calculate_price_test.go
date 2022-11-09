package usecase

import (
	"database/sql"

	"github.com/stretchr/testify/suite"
	"github.com/thiagorcode/gointensivo/internal/order/infra/database"

	_ "github.com/mattn/go-sqlite3"
)

type CalculatePriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

func (suite *CalculatePriceUseCaseTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	// create table orders
	_, err = db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}
