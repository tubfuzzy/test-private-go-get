package service

import "errors"

func GetB() (Service, error) {
	return Service{}, errors.New("error: error example")
}
