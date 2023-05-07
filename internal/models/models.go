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
	Color          uint32    `json:"color"`
	BodyType       uint32    `json:"body_type"`
	DriveGear      uint32    `json:"drive_gear"`
	GearBox        uint32    `json:"gear_box"`
	EngineType     uint32    `json:"engine_type"`
	EngineCapacity uint32    `json:"engine_capacity"`
	EnginePower    uint32    `json:"engine_power"`
	Description    string    `json:"description,omitempty"`
}
