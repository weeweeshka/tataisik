package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/weeweeshka/tataisk/internal/domain/models"
	"go.uber.org/zap"
)

type Storage struct {
	db *pgxpool.Pool
}

func configurationPool(config *pgxpool.Config) {
	config.MaxConns = int32(20)
	config.MinConns = int32(5)
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = time.Minute
	config.ConnConfig.ConnectTimeout = 5 * time.Second
}

func NewStorage(connString string, logr *zap.Logger) (*Storage, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	configurationPool(config)

	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}
	logr.Info("connected to database", zap.String("path to db", connString))
	return &Storage{db: dbPool}, nil

}

func (s *Storage) CreateFilmDB(ctx context.Context, data models.FilmData) (int32, error) {
	var filmID int32

	err := s.db.QueryRow(ctx, `INSERT INTO tataisk(
	                      title, year_of_prod, imdb, description, country,
	                      genre, film_director, screenwriter, budget, collection)
	                      VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		data.Title, data.YearOfProd, data.Imdb, data.Description, data.Country,
		data.Genre, data.FilmDirector, data.Screenwriter, data.Budget, data.Collection).Scan(&filmID)
	if err != nil {
		return 0, fmt.Errorf("error creating film: %w", err)
	}

	return filmID, nil
}

func (s *Storage) ReadFilmDB(ctx context.Context, filmID int32) (models.FilmData, error) {

	filmData := models.FilmData{}

	err := s.db.QueryRow(ctx, `SELECT id, title, year_of_prod, imdb, description, country,
						 genre, film_director, screenwriter, budget, collection FROM tataisk WHERE id = $1`, filmID,
	).Scan(&filmData.Id,
		&filmData.Title,
		&filmData.YearOfProd,
		&filmData.Imdb,
		&filmData.Description,
		&filmData.Country,
		&filmData.Genre,
		&filmData.FilmDirector,
		&filmData.Screenwriter,
		&filmData.Budget,
		&filmData.Collection)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.FilmData{}, errors.New("film not found")
		}
		return models.FilmData{}, fmt.Errorf("error reading film: %w", err)
	}

	return filmData, nil
}

func (s *Storage) UpdateFilmDB(ctx context.Context, filmID int32, data models.FilmData) (bool, error) {

	s.db.QueryRow(ctx, `UPDATE tataisk SET title = $1, year_of_prod = $2, imdb = $3, description = $4, country = $5,
						 genre = $6, film_director = $7, screenwriter = $8, budget = $9, collection = $10 WHERE id = $11`,
		data.Title, data.YearOfProd, data.Imdb, data.Description, data.Country, data.Genre, data.FilmDirector, data.Screenwriter, data.Budget, data.Collection, filmID)

	return true, nil
}

func (s *Storage) DeleteFilmDB(ctx context.Context, filmID int32) (bool, error) {

	s.db.QueryRow(ctx, `DELETE FROM tataisk WHERE id = $1`, filmID)
	return true, nil
}
