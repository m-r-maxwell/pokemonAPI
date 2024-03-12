package pokemon

import "ketchum/models"

type IRepo interface {
	GetPokemonByName(name string) (models.Pokemon, error)
	GetPokemonByType(pokemonType string) ([]models.Pokemon, error)
	GetPokemonByPokedexNumber(id int) (models.Pokemon, error)
	GetPokemonList() ([]models.Pokemon, error)
	GetPokemonAbilities() ([]models.Ability, error)
	GetPokemonAbilityByName(name string) (models.Ability, error)
	CreatePokemon(pokemon models.Pokemon) error
}
