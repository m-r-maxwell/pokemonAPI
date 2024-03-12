package pokemon

import (
	"ketchum/models"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewPokemonRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// GetPokemon returns a pokemon by its name
func (r *Repo) GetPokemonByName(name string) (models.Pokemon, error) {
	sql := `
	SELECT
	p.name,
	p.pokedex_number,
	p.type1,
	p.type2,
	p.stats::jsonb
	FROM pokemon.pokemon p
	WHERE name = $1`
	pokemon := models.Pokemon{}
	err := r.db.Get(&pokemon, sql, name)
	return pokemon, err
}

func (r *Repo) GetPokemonByType(pokemonType string) ([]models.Pokemon, error) {
	sql := `
	SELECT
	p.name,
	p.pokedex_number,
	p.type1,
	p.type2,
	p.stats
	FROM pokemon.pokemon p
	WHERE type1 = $1 OR type2 = $1
	`
	pokemon := make([]models.Pokemon, 0)
	err := r.db.Select(&pokemon, sql, pokemonType)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

func (r *Repo) GetPokemonByPokedexNumber(id int) (models.Pokemon, error) {
	sql := `
	SELECT 
	p.name,
	p.pokedex_number,
	p.type1,
	p.type2,
	p.stats
	FROM pokemon.pokemon p
	WHERE pokedex_number = $1`
	pokemon := models.Pokemon{}
	err := r.db.Get(&pokemon, sql, id)
	if err != nil {
		return models.Pokemon{}, err
	}
	return pokemon, nil
}

func (r *Repo) GetPokemonList() ([]models.Pokemon, error) {
	sql := `
	SELECT 
	p.name,
	p.pokedex_number,
	p.type1,
	p.type2,
	p.stats

	from pokemon.pokemon p 
	`
	pokemon := []models.Pokemon{}
	err := r.db.Select(&pokemon, sql)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

func (r *Repo) GetPokemonAbilities() ([]models.Ability, error) {
	sql := `SELECT * FROM pokemon.abilities`
	abilities := []models.Ability{}
	err := r.db.Select(&abilities, sql)
	if err != nil {
		return nil, err
	}
	return abilities, nil
}

func (r *Repo) GetPokemonAbilityByName(name string) (models.Ability, error) {
	sql := `SELECT * FROM pokemon.abilities WHERE ability_name = $1`
	ability := models.Ability{}
	err := r.db.Get(&ability, sql, name)
	if err != nil {
		return models.Ability{}, err
	}
	return ability, nil
}

func (r *Repo) CreatePokemon(pokemon models.Pokemon) error {
	return nil
}
