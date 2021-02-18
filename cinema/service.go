package cinema

import "github.com/kalmad99/Ethio_Cinema_Final_REST/model"

type CinemaRepository interface {
	Cinemas() ([]model.Cinema, []error)
	Cinema(id uint) (*model.Cinema, []error)
	StoreCinema(cinema *model.Cinema) (*model.Cinema, []error)
	UpdateCinema(cinema *model.Cinema) (*model.Cinema, []error)
	DeleteCinema(id uint) (*model.Cinema, []error)
	CinemaExists(cinemaName string) bool
}