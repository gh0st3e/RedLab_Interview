package handler

import (
	"context"
	"github.com/gh0st3e/RedLab_Interview/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductService interface {
	SaveProduct(ctx context.Context, product entity.Product, userID int) error
	RetrieveProduct(ctx context.Context, barcode string, userID int) (*entity.Product, error)
	DeleteProduct(ctx context.Context, barcode string, userID int) error
	RetrieveProductsByUserID(ctx context.Context, userID int) ([]entity.Product, error)
}

func (h *Handler) SaveProduct(c *gin.Context) {
	var product entity.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		h.logger.Errorf("[SaveProduct] Error while parse product struct: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data, use correct",
		})
		return
	}

	ctxID, ok := c.Get(userIDCtx)
	if !ok {
		h.logger.Errorf("[SaveProduct] Couldn't get user id from context")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get your credentials, try later",
		})
		return
	}

	userID, ok := ctxID.(int)
	if !ok {
		h.logger.Errorf("[SaveProduct] Couldn't parse user id to int")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get your credentials, try later",
		})
		return
	}

	err = h.service.SaveProduct(c, product, userID)
	if err != nil {
		h.logger.Errorf("[SaveProduct] Couldnt save product: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "Successfully Saved",
	})
}

func (h *Handler) RetrieveProduct(c *gin.Context) {
	productID := c.Param("barcode")
	if productID == "" {
		h.logger.Errorf("[RetrieveProduct] Couldn't get product id from query param")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get product id from query param",
		})
		return
	}

	ctxID, ok := c.Get(userIDCtx)
	if !ok {
		h.logger.Errorf("[RetrieveProduct] Couldn't get user id from context")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get your credentials, try later",
		})
		return
	}

	userID, ok := ctxID.(int)
	if !ok {
		h.logger.Errorf("[RetrieveProduct] Couldn't parse user id to int")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get your credentials, try later",
		})
		return
	}

	product, err := h.service.RetrieveProduct(c, productID, userID)
	if err != nil {
		h.logger.Errorf("[RetrieveProduct] Error while retrieving product: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	productID := c.Param("barcode")
	if productID == "" {
		h.logger.Errorf("[DeleteProduct] Couldn't get product id from query param")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get product id from query param",
		})
		return
	}

	ctxID, ok := c.Get(userIDCtx)
	if !ok {
		h.logger.Errorf("[DeleteProduct] Couldn't get user id from context")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get your credentials, try later",
		})
		return
	}

	userID, ok := ctxID.(int)
	if !ok {
		h.logger.Errorf("[DeleteProduct] Couldn't parse user id to int")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get your credentials, try later",
		})
		return
	}

	err := h.service.DeleteProduct(c, productID, userID)
	if err != nil {
		h.logger.Errorf("[DeleteProduct] Error while deleting product: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "Successfully Deleted",
	})
}

func (h *Handler) RetrieveProductsByUserID(c *gin.Context) {
	ctxID, ok := c.Get(userIDCtx)
	if !ok {
		h.logger.Errorf("[RetrieveProductsByUserID] Couldn't get user id from context")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get your credentials, try later",
		})
		return
	}

	userID, ok := ctxID.(int)
	if !ok {
		h.logger.Errorf("[RetrieveProductsByUserID] Couldn't parse user id to int")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get your credentials, try later",
		})
		return
	}

	products, err := h.service.RetrieveProductsByUserID(c, userID)
	if err != nil {
		h.logger.Errorf("[RetrieveProductsByUserID] Couldn't get user products.sql: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
