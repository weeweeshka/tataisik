package tataisk

import (
	"context"
	"fmt"
	"github.com/weeweeshka/tataisk/internal/domain/models"
	"go.uber.org/zap"
)

type Tataisk struct {
	logr *zap.Logger
	repo FilmRepository
}

type FilmRepository interface {
	CreateFilmDB(ctx context.Context, data models.FilmData) (int32, error)
	ReadFilmDB(ctx context.Context, id int32) (models.FilmData, error)
	UpdateFilmDB(ctx context.Context, id int32, data models.FilmData) (bool, error)
	DeleteFilmDB(ctx context.Context, id int32) (bool, error)
}

func New(logger *zap.Logger, repo FilmRepository) *Tataisk {
	return &Tataisk{logr: logger, repo: repo}
}

func (t *Tataisk) CreateFilm(ctx context.Context, data models.FilmData) (int32, error) {

	filmId, err := t.repo.CreateFilmDB(ctx, data)
	if err != nil {
		return 0, fmt.Errorf("error creating film: %w", err)
	}

	t.logr.Info("Film created!")

	return filmId, nil
}

func (t *Tataisk) ReadFilm(ctx context.Context, id int32) (models.FilmData, error) {

	film, err := t.repo.ReadFilmDB(ctx, id)
	if err != nil {
		return models.FilmData{}, fmt.Errorf("error reading film: %w", err)
	}
	t.logr.Info("Film read!")

	return film, nil
}

func (t *Tataisk) UpdateFilm(ctx context.Context, id int32, data models.FilmData) (bool, error) {

	success, err := t.repo.UpdateFilmDB(ctx, id, data)
	if err != nil {
		return success, fmt.Errorf("error updating film: %w", err)
	}
	t.logr.Info("Film updated!")

	return success, nil
}

func (t *Tataisk) DeleteFilm(ctx context.Context, id int32) (bool, error) {

	success, err := t.repo.DeleteFilmDB(ctx, id)
	if err != nil {
		return success, fmt.Errorf("error deleting film: %w", err)
	}
	t.logr.Info("Film deleted!")

	return success, nil
}
