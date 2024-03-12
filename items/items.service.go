package items

import (
	"ketchum/helpers"
	"ketchum/models"
	"net/http"
)

type ItemService struct {
	repo IItemRepo
}

func NewItemService(repo IItemRepo) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (is *ItemService) RegisterHandlers(mux *http.ServeMux) {
	mux.Handle("GET /items", is.GetAllItems())
	mux.Handle("POST /items", is.GetItemByName())
	mux.Handle("POST /items/price", is.GetItemByPriceRange())
	mux.Handle("POST /items/new", is.CreateItem())
}

// GetAllItems returns a list of all items
func (is *ItemService) GetAllItems() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// swagger:route GET /items items getAllItems
		// Returns a list of all items
		// Consumes:
		// - application/json
		// Produces:
		// - application/json
		// Responses:
		//   200: []Item
		//   500: error
		items, err := is.repo.GetAllItems()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, items)
	})
}

// GetItemByName returns an item by its name
func (is *ItemService) GetItemByName() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// swagger:route POST /items items getItemByName
		// Returns an item by its name
		// Consumes:
		// - application/json
		// Produces:
		// - application/json
		// Params:
		// - name: item
		//   in: body
		//   description: The name of the item
		//   required: true
		// Responses:
		//   200: Item
		//   500: error
		var it models.Item
		err := helpers.DecodeBody(r, &it)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		item, err := is.repo.GetItemByName(it.Name)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, item)
	})
}

// GetItemByPriceRange returns a list of items within a price range
func (is *ItemService) GetItemByPriceRange() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// swagger:route POST /items/price items getItemByPriceRange
		// Returns a list of items within a price range
		// Consumes:
		// - application/json
		// Produces:
		// - application/json
		// Params:
		// - name: minPrice
		//   in: body
		//   description: The price range
		//   required: true
		// - name: maxPrice
		//   in: body
		//   description: The price range
		//   required: true
		// Responses:
		//   200: []Item
		//   500: error
		var pr models.PriceRange
		err := helpers.DecodeBody(r, &pr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items, err := is.repo.GetItemByPriceRange(pr.MinPrice, pr.MaxPrice)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, items)
	})
}

func (is *ItemService) CreateItem() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var item models.Item
		err := helpers.DecodeBody(r, &item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = is.repo.CreateItem(item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		helpers.JSONResponse(w, item)
	})
}
