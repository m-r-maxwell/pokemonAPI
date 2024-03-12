package models

import "encoding/json"

// swagger:model Pokemon
type Pokemon struct {
	PokedexNumber int             `json:"pokedexNumber" db:"pokedex_number"`
	Name          string          `json:"name" db:"name"`
	Type1         string          `json:"type1" db:"type1"`
	Type2         string          `json:"type2" db:"type2"`
	Stats         json.RawMessage `json:"stats" db:"stats"`
}

// swagger:model Ability
type Ability struct {
	AbilityID   int    `db:"ability_id"`
	Name        string `json:"name" db:"ability_name"`
	Description string `json:"description" db:"ability_description"`
}

// swagger:model Move
type Move struct {
	Id              int    `db:"move_id"`
	Name            string `json:"name" db:"move_name"`
	Category        string `json:"category" db:"move_category"`
	Type            string `json:"type" db:"move_type"`
	Power           int    `json:"power" db:"move_power"`
	Accuracy        int    `json:"accuracy" db:"move_accuracy"`
	PP              int    `json:"pp" db:"move_pp"`
	PriorityLevel   int    `json:"priorityLevel" db:"move_priority_level"`
	Priority        bool   `json:"priority" db:"move_priority"`
	Contact         bool   `json:"contact" db:"move_contact"`
	SecondaryEffect string `json:"secondaryEffect" db:"move_effect"`
	SecondaryChance int    `json:"secondaryChance" db:"move_effect_chance"`
}

// swagger:model Item
type Item struct {
	Name        string `json:"name" db:"item_name"`
	Price       int    `json:"price" db:"item_price"`
	Description string `json:"description" db:"item_description"`
}

// swagger:model PriceRange
type PriceRange struct {
	MinPrice int `json:"minPrice"`
	MaxPrice int `json:"maxPrice"`
}
