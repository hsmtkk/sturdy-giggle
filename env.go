package main

import (
	"fmt"
	"os"
	"strconv"
)

func loadEnvString(name string) (string, error) {
	s := os.Getenv(name)
	if s == "" {
		return "", fmt.Errorf("%s env var is not defined", name)
	}
	return s, nil
}

func loadEnvInt(name string) (int, error) {
	s, err := loadEnvString(name)
	if err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s as int: %w", s, err)
	}
	return n, nil
}

type postgresConfig struct {
	postgresHost     string
	postgresPort     int
	postgresDB       string
	postgresUser     string
	postgresPassword string
}

func loadPostgresConfig() (postgresConfig, error) {
	pgConfig := postgresConfig{}
	var err error
	pgConfig.postgresDB, err = loadEnvString("POSTGRES_DB")
	if err != nil {
		return pgConfig, err
	}
	pgConfig.postgresHost, err = loadEnvString("POSTGRES_HOST")
	if err != nil {
		return pgConfig, err
	}
	pgConfig.postgresPassword, err = loadEnvString("POSTGRES_PASSWORD")
	if err != nil {
		return pgConfig, err
	}
	pgConfig.postgresPort, err = loadEnvInt("POSTGRES_PORT")
	if err != nil {
		return pgConfig, err
	}
	pgConfig.postgresUser, err = loadEnvString("POSTGRES_USER")
	if err != nil {
		return pgConfig, err
	}
	return pgConfig, nil
}
