package controller

import (
	"database/sql"
	"errors"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//usecase
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertProduct, err := p.productUseCase.CreateProducts(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um numero.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if err == nil && product == nil {
		response := model.Response{
			Message: "Produto não foi encontrado na base de dados.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Id do produto não pode ser nulo.",
		})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Id do produto precisa ser um numero.",
		})
		return
	}

	err = p.productUseCase.DeleteProductById(productId)

	if err != nil {

		// Produto não encontrado
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, model.Response{
				Message: "Produto não foi encontrado na base de dados.",
			})
			return
		}

		// Outro erro
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Erro interno ao deletar produto.",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Produto deletado com sucesso.",
	})
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	id := ctx.Param("productId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Id do produto não pode ser nulo.",
		})
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Id do produto precisa ser um numero.",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updateProduct, err := p.productUseCase.UpdateteProduct(product, productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, updateProduct)
}
