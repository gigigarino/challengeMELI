package repository

import "github.com/gigigarino/challengeMELI2/internal/domain"

//definimos estructura
type itemRepository struct {
	articulos []domain.Item
}

//definimos nuevafunciom
func NewItemRepository() domain.ItemRepository {
	return &itemRepository{}
}


func (r *itemRepository) GetAllItems(){
	return r.articulos
}

func (r *itemRepository) GetItemById(id int) *domain.item{
	for _, item := range r.articulos {
		if item.ID == id {
			return item
		}
	}
	return nil
}

func (r * itemRepository) AddItem (item domain.Item) *domain.Item{
	r.articulos = append(r.articulos, item)
	return &item
}