package items

import "ketchum/models"

type IItemRepo interface {
	GetItemByName(name string) (models.Item, error)
	GetItemByPriceRange(min, max int) ([]models.Item, error)
	GetAllItems() ([]models.Item, error)
	CreateItem(item models.Item) error
}
