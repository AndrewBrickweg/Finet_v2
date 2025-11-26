package database

import (
	"context"
	"log"
)

type StockData struct {
	Ticker    string  `json:"ticker"`
	Date      string  `json:"date"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	AdjClose  float64 `json:"adj_close"`
	Volume    int64   `json:"volume"`
	Dividend  float64 `json:"dividend"`
}

type StockDB struct {
	DBService *DBService
}

func NewStockDB(ctx context.Context, dataSourceName string) (*StockDB, error) {
	base, err:= NewDBService(ctx, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &StockDB{DBService: base}, nil	
}

func (s *StockDB) Close() error {
	log.Println("Closing stock database connection.")
	return s.DBService.Close()
}
