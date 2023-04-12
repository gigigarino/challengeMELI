package usecase

import (
	"github.com/gigigarino/challengeMELI"
)
/*
en esta capa definimos una estructura y una interface con los metodos o casos de uso
getallbooks
getbookporid
addbook
y retornamos el newbook usecase 

se definen todas las funciones 
*/

//definimos una interface 

type ItemUsecase interface {
	//metodos 
	GetAllItems() []domain.Item
	GetItemById(id int) *domain.Item
	AddItem (item domain.Item)(*domain.Item, error)
}


//definimos una estructura
type bookUsecase struct {
	repo domain.ItemRepository
}

//funcion new
func NewBookUsecase(repo domain.ItemRepository) ItemUsecase {
	return &ItemUsecase{
		repo: repo,
	}
}