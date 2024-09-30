package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mubinabd/modestyMart/internal/pkg/config"
	"github.com/Mubinabd/modestyMart/internal/storage"
)

type Storage struct {
	Db        *sql.DB
	OrderS    storage.OrderI
	ProductS  storage.ProductI
	PaymentS  storage.PaymentI
	CategoryS storage.CategoryI
	CartS     storage.CartI
	AuthS     storage.AuthI
	UserS     storage.UserI
}

func New(cfg *config.Config) (*Storage, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Println("can't connect to db: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{
		Db:        db,
		OrderS:    NewOrderRepo(db),
		ProductS:  NewProductRepo(db),
		PaymentS:  NewPaymentRepo(db),
		CategoryS: NewCategoryRepo(db),
		AuthS:     NewAuthRepo(db),
		UserS:     NewUserRepo(db),
		CartS:     NewCartRepo(db),
	}, nil
}

func (s *Storage) Order() storage.OrderI {
	if s.OrderS == nil {
		s.OrderS = NewOrderRepo(s.Db)
	}
	return s.OrderS
}

func (s *Storage) Product() storage.ProductI {
	if s.ProductS == nil {
		s.ProductS = NewProductRepo(s.Db)
	}
	return s.ProductS
}
func (s *Storage) Payment() storage.PaymentI {
	if s.PaymentS == nil {
		s.PaymentS = NewPaymentRepo(s.Db)
	}
	return s.PaymentS
}

func (s *Storage) Category() storage.CategoryI {
	if s.CategoryS == nil {
		s.CategoryS = NewCategoryRepo(s.Db)
	}
	return s.CategoryS
}

func (s *Storage) Cart() storage.CartI {
	if s.CartS == nil {
		s.CartS = NewCartRepo(s.Db)
	}
	return s.CartS
}

func (s *Storage) Auth() storage.AuthI {
	if s.AuthS == nil {
		s.AuthS = NewAuthRepo(s.Db)
	}
	return s.AuthS
}

func (s *Storage) User() storage.UserI {
	if s.UserS == nil {
		s.UserS = NewUserRepo(s.Db)
	}
	return s.UserS
}
