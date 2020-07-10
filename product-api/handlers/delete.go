package handlers

import (
	"microservices/product-api/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes the product with the given id
// responses:
// 201: noContent
// 404: errorResponse
// 501: errorResponse

// DeleteProduct deletes a product
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("Handle DELETE Product with id: ", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] Record not found")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] Can't delete product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
