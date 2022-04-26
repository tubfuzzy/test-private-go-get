package service

import "errors"

func GetA() (Service, error) {
	return Service{Data: "example"}, nil
}

func GetC() (Service, error) {
	return Service{}, errors.New("error: error example C")
}
