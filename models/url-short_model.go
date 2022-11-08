package models

type URLShort struct {
	Id  string `json:"Id,omitempty" validate:"required"`
	Url string `json:"url,omitempty" validate:"required"`
}