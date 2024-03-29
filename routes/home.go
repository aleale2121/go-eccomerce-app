package routes

import (
	"fmt"
	"net/http"

	"github.com/aleale2121/go-eccomerce-app/models"
	"github.com/aleale2121/go-eccomerce-app/utils"
)

func  (rs Routes) homeGetHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := rs.store.GetCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	products, err := rs.store.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	utils.ExecuteTemplate(w, "home.html", struct {
		Categories []models.Category
		Products   []models.Product
	}{
		Categories: categories,
		Products:   products,
	})
}

func  (rs Routes) homePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=UTF-8")
	r.ParseForm()
	search := r.PostForm.Get("search")
	products, err := rs.store.SearchProducts(search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	count := len(products)
	var html string = ""
	if count > 0 {
		html += "<table class='table table-bordered'>"
		html += "<th> Id </th> <th> Category </th> <th> Nome </th> <th> Preço </th> <th> Amount </th> <th> Valor total</th>"
		for _, p := range products {
			html += "<tr>"
			html += fmt.Sprintf("<td> %d </td> <td> %s </td> <td> %s </td> <td> %.2f R$ </td> <td> %d </td> <td> %.2f </td>", p.Id,
				p.Category.Description, p.Name, p.Price, p.Quantity, p.Amount)
			html += "</tr>"
		}
		html += "</table>"
	} else {
		html += fmt.Sprintf(`<p class='alert alert-info'> nothing found with <code>"<strong> %s </strong> </code>"</p>`, search)
	}

	w.Write([]byte(html))
}
