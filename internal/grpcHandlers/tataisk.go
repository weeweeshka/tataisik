package grpcHandlers

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/weeweeshka/tataisk/internal/domain/models"
	pb "github.com/weeweeshka/tataisk_proto/gen/go/tataisk"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//TODO исправить на float32 в контракте

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

type Tataisk interface {
	CreateFilm(
		ctx context.Context,
		data models.FilmData,
	) (int32, error)

	ReadFilm(
		ctx context.Context,
		id int32,
	) (models.FilmData, error)

	UpdateFilm(
		ctx context.Context,
		id int32,
		data models.FilmData,
	) (bool, error)

	DeleteFilm(
		ctx context.Context,
		id int32,
	) (bool, error)
}

type serverAPI struct {
	pb.UnimplementedTataiskServer
	api Tataisk
}

func RegisterNewServer(gRPCServer *grpc.Server, api Tataisk) {
	pb.RegisterTataiskServer(gRPCServer, &serverAPI{api: api})
	InitValidator()
}

func (s *serverAPI) CreateFilm(ctx context.Context, req *pb.CreateFilmRequest) (*pb.CreateFilmResponse, error) {

	filmData := models.FilmData{
		Title:        req.GetTitle(),
		YearOfProd:   req.GetYearOfProd(),
		Imdb:         req.GetImdb(),
		Description:  req.GetDescription(),
		Country:      req.GetCountry(),
		Genre:        req.GetGenre(),
		FilmDirector: req.GetFilmDirector(),
		Screenwriter: req.GetScreenwriter(),
		Budget:       req.GetBudget(),
		Collection:   req.GetCollection(),
	}

	err := validate.Struct(filmData)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := s.api.CreateFilm(ctx, filmData)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateFilmResponse{Id: id}, nil
}

func (s *serverAPI) ReadFilm(ctx context.Context, req *pb.ReadFilmRequest) (*pb.ReadFilmResponse, error) {
	var film models.FilmData

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	film, err := s.api.ReadFilm(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.ReadFilmResponse{
		Id:           req.GetId(),
		Title:        film.Title,
		YearOfProd:   film.YearOfProd,
		Imdb:         film.Imdb,
		Description:  film.Description,
		Country:      film.Country,
		Genre:        film.Genre,
		FilmDirector: film.FilmDirector,
		Screenwriter: film.Screenwriter,
		Budget:       film.Budget,
		Collection:   film.Collection,
	}, nil

}

func (s *serverAPI) UpdateFilm(ctx context.Context, req *pb.UpdateFilmRequest) (*pb.UpdateFilmResponse, error) {

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	filmData := models.FilmData{
		Title:        req.GetTitle(),
		YearOfProd:   req.GetYearOfProd(),
		Imdb:         req.GetImdb(),
		Description:  req.GetDescription(),
		Country:      req.GetCountry(),
		Genre:        req.GetGenre(),
		FilmDirector: req.GetFilmDirector(),
		Screenwriter: req.GetScreenwriter(),
		Budget:       req.GetBudget(),
		Collection:   req.GetCollection(),
	}

	success, err := s.api.UpdateFilm(ctx, req.GetId(), filmData)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateFilmResponse{Success: success}, nil
}

func (s *serverAPI) DeleteFilm(ctx context.Context, req *pb.DeleteFilmRequest) (*pb.DeleteFilmResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	success, err := s.api.DeleteFilm(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.DeleteFilmResponse{Success: success}, nil
}
