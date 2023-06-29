package routes

import (
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/aleale2121/go-eccomerce-app/models"
	"github.com/aleale2121/go-eccomerce-app/sessions"
	"github.com/aleale2121/go-eccomerce-app/utils"
)

var (
	ErrPriceValue          = errors.New("input error: \"price\" invalid")
	ErrQuantityValue       = errors.New("input error: \"amount\" invalid")
	ErrRequiredProductName = errors.New("required product name")
)

func (rs Routes) productGetHandler(w http.ResponseWriter, r *http.Request) {
	products, err := rs.store.GetProducts()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	total := int64(len(products))
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "product.html", struct {
		Total    int64
		Products []models.Product
		Alert    utils.Alert
	}{
		Total:    total,
		Products: products,
		Alert:    utils.NewAlert(message, alert),
	})
}

func  (rs Routes) productCreateGetHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := rs.store.GetCategories()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "product_create.html", struct {
		Categories []models.Category
		Alert      utils.Alert
	}{
		Categories: categories,
		Alert:      utils.NewAlert(message, alert),
	})
}

func  (rs Routes) productCreatePostHandler(w http.ResponseWriter, r *http.Request) {
	product, err := rs.verifyInputsProduct(r)
	if err != nil {
		sessions.Message(fmt.Sprintf("%s", err), "danger", r, w)
		http.Redirect(w, r, "/product-create", http.StatusSeeOther)
		return
	}
	_, err = rs.store.NewProduct(product)
	if err != nil {
		log.Println(err)
		utils.InternalServerError(w)
		return
	}
	sessions.Message("New product added", "success", r, w)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func  (rs Routes) verifyInputsProduct(r *http.Request) (models.Product, error) {
	r.ParseForm()
	var err error = nil
	var product models.Product
	product.Id, _ = strconv.ParseUint(r.PostForm.Get("id"), 10, 64)
	product.Name = html.EscapeString(r.PostForm.Get("name"))
	if models.IsEmpty(product.Name) {
		return product, ErrRequiredProductName
	}
	if !models.Max(product.Name, 255) {
		return product, models.ErrMaxLimit
	}
	product.Price, err = strconv.ParseFloat(r.PostForm.Get("price"), 64)
	if err != nil {
		return product, ErrPriceValue
	}
	product.Quantity, err = strconv.Atoi(r.PostForm.Get("quantity"))
	if err != nil {
		return product, ErrQuantityValue
	}
	product.Amount = float64(product.Quantity) * product.Price
	product.Category.Id, _ = strconv.Atoi(r.PostForm.Get("category"))
	return product, nil
}

func (rs Routes) productEditGetHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	productId, _ := strconv.ParseUint(keys.Get("productId"), 10, 64)
	product, err := rs.store.GetProductById(productId)
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	categories, err := rs.store.GetCategories()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	priceFormat := product.PriceToString()
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "product_edit.html", struct {
		Categories  []models.Category
		Product     models.Product
		PriceFormat string
		Alert       utils.Alert
	}{
		Categories:  categories,
		Product:     product,
		PriceFormat: priceFormat,
		Alert:       utils.NewAlert(message, alert),
	})
}

func  (rs Routes) productEditPostHandler(w http.ResponseWriter, r *http.Request) {
	product, err := rs.verifyInputsProduct(r)
	if err != nil {
		sessions.Message(fmt.Sprintf("%s", err), "danger", r, w)
		http.Redirect(w, r, fmt.Sprintf("product-edit?productId=%d", product.Id), http.StatusSeeOther)
		return
	}
	rows, err := rs.store.UpdateProduct(product)
	if err != nil {
		log.Println(err)
		utils.InternalServerError(w)
		return
	}
	sessions.Message(fmt.Sprintf("%d product has been updated successfully!", rows), "info", r, w)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func  (rs Routes) productDeleteGetHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	id, _ := strconv.ParseUint(keys.Get("productId"), 10, 64)
	ok, _ := strconv.ParseBool(keys.Get("confirm"))
	if !ok {
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}
	rows, err := rs.store.DeleteProduct(id)
	if err != nil {
		log.Println(err)
		utils.InternalServerError(w)
		return
	}
	sessions.Message(fmt.Sprintf("%d product has been permanently deleted.", rows), "warning", r, w)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
