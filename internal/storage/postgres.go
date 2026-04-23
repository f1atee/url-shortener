package storage

import (
	"database/sql"
	"fmt"
)

var ErrNotFound = fmt.Errorf("url not found")

type Storage struct {
	db *sql.DB
}

func New(dsn string) (*Storage, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(code, url string) error {
	_, err := s.db.Exec(
		"INSERT INTO urls (code, url) VALUES ($1, $2)",
		code, url,
	)
	if err != nil {
		return fmt.Errorf("insert url: %w", err)
	}
	return nil
}

func (s *Storage) GetURL(code string) (string, error) {
	var url string
	err := s.db.QueryRow(
		"SELECT url FROM urls WHERE code = $1", code,
	).Scan(&url)

	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("query url: %w", err)
	}
	return url, nil
}

func (s *Storage) DeleteURL(code string) error {
	result, err := s.db.Exec("DELETE FROM urls WHERE code = $1", code)
	if err != nil {
		return fmt.Errorf("delete url: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
