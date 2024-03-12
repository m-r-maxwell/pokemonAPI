package items

import (
	"fmt"
	"ketchum/models"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// GetAllItems returns a list of all items
func (r *Repo) GetAllItems() ([]models.Item, error) {
	sql := `
	SELECT
	i.item_name,
	i.item_description,
	i.item_price
	FROM items.items i
	`
	items := make([]models.Item, 0)
	err := r.db.Select(&items, sql)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// GetItemByName returns an item by its name
func (r *Repo) GetItemByName(name string) (models.Item, error) {
	sql := `
	SELECT
	i.item_name,
	i.item_description,
	i.item_price
	FROM items.items i
	WHERE item_name = $1
	`
	item := models.Item{}
	err := r.db.Get(&item, sql, name)
	return item, err
}

// GetItemByPriceRange returns a list of items within a price range
func (r *Repo) GetItemByPriceRange(min, max int) ([]models.Item, error) {
	sql := `
	SELECT
	i.item_name,
	i.item_description,
	i.item_price
	FROM items.items i
	WHERE item_price BETWEEN $1 AND $2
	`
	items := make([]models.Item, 0)
	err := r.db.Select(&items, sql, min, max)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repo) CreateItem(item models.Item) error {
	sql := `
	INSERT INTO items.items (item_name, item_description, item_price)
	VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(sql, item.Name, item.Description, item.Price)
	fmt.Println(err)
	return err
}
