package moves

import (
	"ketchum/helpers"
	"ketchum/models"
	"net/http"
)

type MoveService struct {
	repo IMoveRepo
}

func NewMoveService(repo IMoveRepo) *MoveService {
	return &MoveService{}
}

func (ms *MoveService) RegisterHandlers(mux *http.ServeMux) {
	mux.Handle("GET /moves", ms.GetAllMoves())
	mux.Handle("POST /moves", ms.GetMoveByName())
	mux.Handle("POST /moves/type/{type}", ms.GetMoveByType())
	mux.Handle("POST /moves/new", ms.CreateMove())
}

// GetAllMoves returns a list of all moves
func (ms *MoveService) GetAllMoves() http.HandlerFunc {
	// swagger:route GET /moves moves getAllMoves
	// Returns a list of all moves
	// Consumes:
	// - application/json
	// Produces:
	// - application/json
	// Responses:
	//   200: []Move
	//   500: error
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		moves, err := ms.repo.GetAllMoves()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, moves)
	})
}

// GetMoveByName returns a move by its name
func (ms *MoveService) GetMoveByName() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// swagger:route POST /moves moves getMoveByName
		// Returns a move by its name
		// Consumes:
		// - application/json
		// Produces:
		// - application/json
		// Params:
		// - name: move
		//   in: body
		//   description: The name of the move
		//   required: true
		// Responses:
		//   200: Move
		//   500: error
		var mv models.Move
		err := helpers.DecodeBody(r, &mv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		move, err := ms.repo.GetMoveByName(mv.Name)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, move)
	})
}

// GetMoveByType returns a list of moves by type
func (ms *MoveService) GetMoveByType() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// swagger:route POST /moves/type/{type} moves getMoveByType
		// Returns a list of moves by type
		// Consumes:
		// - application/json
		// Produces:
		// - application/json
		// Params:
		// - name: type
		//   in: path
		//   description: The type of the move
		//   required: true
		// Responses:
		//   200: []Move
		//   500: error
		mt := r.PathValue("type")
		moves, err := ms.repo.GetMoveByType(mt)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, moves)
	})
}

// CreateMove creates a new move
func (ms *MoveService) CreateMove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// swagger:route POST /moves/new moves createMove
		// Creates a new move
		// Consumes:
		// - application/json
		// Produces:
		// - application/json
		// Params:
		// - name: move
		//   in: body
		//   description: The move to create
		//   required: true
		// Responses:
		//   200: Move
		//   500: error
		var mv models.Move
		err := helpers.DecodeBody(r, &mv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = ms.repo.CreateMove(mv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, mv)
	})
}
