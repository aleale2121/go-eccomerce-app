package routes

import (
	"net/http"

	"github.com/aleale2121/go-eccomerce-app/models"
	"github.com/aleale2121/go-eccomerce-app/utils"
)
type Routes struct{
	store models.Store
}

func NewRoutes(store models.Store) Routes{
	return Routes{store: store}
}
func (rs Routes) adminGetHandler(w http.ResponseWriter, r *http.Request) {
	products, users, err := rs.LoadData()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	allProducts := int64(len(products))
	allUsers := int64(len(users))
	LastUser := users[len(users)-1]
	LastProduct := products[len(products)-1]
	utils.ExecuteTemplate(w, "admin.html", struct {
		AllProducts int64
		AllUsers    int64
		LastProduct models.Product
		LastUser    models.User
	}{
		AllProducts: allProducts,
		AllUsers:    allUsers,
		LastProduct: LastProduct,
		LastUser:    LastUser,
	})
}

func (rs Routes) LoadData() ([]models.Product, []models.User, error) {
	products, err := rs.store.GetProducts()
	if err != nil {
		return nil, nil, err
	}
	users, err := rs.store.GetUsers()
	if err != nil {
		return nil, nil, err
	}
	return products, users, nil
}
