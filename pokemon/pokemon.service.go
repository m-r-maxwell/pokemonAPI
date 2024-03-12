package pokemon

import (
	"fmt"
	"ketchum/helpers"
	"ketchum/models"
	"net/http"
	"strconv"
)

type PokemonService struct {
	repo IRepo
}

func NewPokemonService(repo IRepo) *PokemonService {
	return &PokemonService{
		repo: repo,
	}
}

func (ps *PokemonService) RegisterHandlers(mux *http.ServeMux) {
	mux.Handle("GET /pokemon", ps.GetPokemonList())
	mux.Handle("POST /pokemon", ps.GetPokemonByName())
	mux.Handle("POST /pokemon/type/{type}", ps.GetPokemonByType())
	mux.Handle("GET /pokemon/{id}", ps.GetPokemonByPokedexNumber())
	mux.Handle("GET /pokemon/abilities", ps.GetPokemonAbilities())
	mux.Handle("POST /pokemon/abilities", ps.GetPokemonAbilityByName())
	mux.Handle("POST /pokemon/new", ps.CreatePokemon())
}

// GetPokemon returns a pokemon by its name
func (ps *PokemonService) GetPokemonByName() http.HandlerFunc {
	// swagger:route POST /pokemon pokemon getPokemonByName
	// Returns a pokemon by its name
	// Consumes:
	// - application/json
	// Produces:
	// - application/json
	// Params:
	// - name: pokemon
	//   in: body
	//   description: The name of the pokemon
	//   required: true
	// Responses:
	//   200: Pokemon
	//   500: error
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var poke models.Pokemon
		err := helpers.DecodeBody(r, &poke)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pokemon, err := ps.repo.GetPokemonByName(poke.Name)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, pokemon)
	})
}

// GetPokemonList returns a list of pokemon up to the 8th generation
func (ps *PokemonService) GetPokemonList() http.HandlerFunc {
	// swagger:route GET /pokemon pokemon getPokemonList
	// Returns a list of pokemon up to the 8th generation
	// Produces:
	// - application/json
	// Responses:
	//   200: []Pokemon
	//   500: error
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokemon, err := ps.repo.GetPokemonList()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, pokemon)
	})
}

func (ps *PokemonService) GetPokemonByType() http.HandlerFunc {
	// swagger:route POST /pokemon/type/{type} pokemon getPokemonByType
	// Returns a list of pokemon by type
	// Consumes:
	// - application/json
	// Produces:
	// - application/json
	// Params:
	// - name: type
	//   in: path
	//   description: The type of the pokemon
	//   required: true
	// Responses:
	//   200: []Pokemon
	//   500: error

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokemonType := r.PathValue("type")
		pokemon, err := ps.repo.GetPokemonByType(pokemonType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, pokemon)
	})
}

func (ps *PokemonService) GetPokemonByPokedexNumber() http.HandlerFunc {
	// swagger:route GET /pokemon/{id} pokemon getPokemonByPokedexNumber
	// Returns a pokemon by its pokedex number
	// Produces:
	// - application/json
	// Params:
	// - name: id
	//   in: path
	//   description: The pokedex number of the pokemon
	//   required: true
	// Responses:
	//   200: Pokemon
	//   500: error
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		pokemonID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		pokemon, err := ps.repo.GetPokemonByPokedexNumber(pokemonID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, pokemon)
	})
}

func (ps *PokemonService) GetPokemonAbilities() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// swagger:route GET /pokemon/abilities pokemon getPokemonAbilities
		// Returns a list of pokemon abilities
		// Produces:
		// - application/json
		// Responses:
		//   200: []Ability
		//   500: error
		abilities, err := ps.repo.GetPokemonAbilities()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, abilities)
	})
}

func (ps *PokemonService) GetPokemonAbilityByName() http.HandlerFunc {
	// swagger:route POST /pokemon/abilities pokemon getPokemonAbilityByName
	// Returns a pokemon ability by its name
	// Consumes:
	// - application/json
	// Produces:
	// - application/json
	// Params:
	// - name: ability
	//   in: body
	//   description: The name of the ability
	//   required: true
	// Responses:
	//   200: Ability
	//   500: error
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var name models.Ability
		err := helpers.DecodeBody(r, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ability, err := ps.repo.GetPokemonAbilityByName(name.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, ability)
	})
}

func (ps *PokemonService) CreatePokemon() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// swagger:route POST /pokemon/new pokemon createPokemon
		// Creates a new pokemon
		// Consumes:
		// - application/json
		// Produces:
		// - application/json
		// Params:
		// - name: pokemon
		//   in: body
		//   description: The pokemon to create
		//   required: true
		// Responses:
		//   200: Pokemon
		//   500: error
		var poke models.Pokemon
		err := helpers.DecodeBody(r, &poke)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = ps.repo.CreatePokemon(poke)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, poke)
	})
}
