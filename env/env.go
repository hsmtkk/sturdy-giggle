package env

import (
	"fmt"
	"os"
	"strconv"
)

type Reader struct {
}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) AsString(name string) (string, error) {
	s := os.Getenv(name)
	if s == "" {
		return "", fmt.Errorf("%s env var is not defined", name)
	}
	return s, nil
}

func (r *Reader) AsInt(name string) (int, error) {
	s, err := r.AsString(name)
	if err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s as int: %w", s, err)
	}
	return n, nil
}

type PostgresConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func (r *Reader) PostgresConfig() (PostgresConfig, error) {
	pgConfig := PostgresConfig{}
	var err error
	pgConfig.Database, err = r.AsString("POSTGRES_DB")
	if err != nil {
		return pgConfig, err
	}
	pgConfig.Host, err = r.AsString("POSTGRES_HOST")
	if err != nil {
		return pgConfig, err
	}
	pgConfig.Password, err = r.AsString("POSTGRES_PASSWORD")
	if err != nil {
		return pgConfig, err
	}
	pgConfig.Port, err = r.AsInt("POSTGRES_PORT")
	if err != nil {
		return pgConfig, err
	}
	pgConfig.User, err = r.AsString("POSTGRES_USER")
	if err != nil {
		return pgConfig, err
	}
	return pgConfig, nil
}
