package models

import "time"

type Publication struct {
	Brand          uint64    `json:"brand"`
	Model          uint64    `json:"model"`
	Vin            string    `json:"vin"`
	ProductionYear time.Time `json:"production_year"`
	Mileage        uint64    `json:"mileage"`
	PicsCount      uint32    `json:"pics_count"`
	OwnerCount     uint32    `json:"owner_count"`
}
