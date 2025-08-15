package models

type FilmData struct {
	Id           int32    `json:"id" validate:"-"`
	Title        string   `json:"title" validate:"required"`
	YearOfProd   uint32   `json:"year_of_prod" validate:"required,min=1900"`
	Imdb         float32  `json:"imdb" validate:"required,min=0,max=10"`
	Description  string   `json:"description" validate:"required"`
	Country      []string `json:"country" validate:"omitempty"`
	Genre        []string `json:"genre" validate:"required"`
	FilmDirector string   `json:"film_director" validate:"required"`
	Screenwriter string   `json:"screenwriter" validate:"required"`
	Budget       int64    `json:"budget" validate:"required"`
	Collection   int64    `json:"collection" validate:"required"`
}
