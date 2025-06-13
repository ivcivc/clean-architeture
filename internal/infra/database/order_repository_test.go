package database

import (
	"database/sql"
	"testing"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"

	"github.com/DATA-DOG/go-sqlmock"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func TestOrderRepository_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewOrderRepository(db)

	rows := sqlmock.NewRows([]string{"id", "price", "tax", "final_price"}).
		AddRow("1", 100.0, 10.0, 110.0).
		AddRow("2", 200.0, 20.0, 220.0)

	mock.ExpectQuery("SELECT id, price, tax, final_price FROM orders").WillReturnRows(rows)

	orders, err := repo.List()
	if err != nil {
		t.Errorf("error was not expected while listing orders: %s", err)
	}

	if len(orders) != 2 {
		t.Errorf("expected 2 orders, but got %d", len(orders))
	}

	if orders[0].ID != "1" || orders[0].Price != 100.0 || orders[0].Tax != 10.0 || orders[0].FinalPrice != 110.0 {
		t.Errorf("first order data mismatch")
	}

	if orders[1].ID != "2" || orders[1].Price != 200.0 || orders[1].Tax != 20.0 || orders[1].FinalPrice != 220.0 {
		t.Errorf("second order data mismatch")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
