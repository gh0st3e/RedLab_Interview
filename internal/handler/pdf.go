package handler

import (
	"github.com/gh0st3e/RedLab_Interview/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PDFService interface {
	LoadPDF(userID int, barcode string) (string, error)
	GeneratePDF(userID int, product entity.Product) (string, error)
}

func (h *Handler) GetPdf(c *gin.Context) {
	productID := c.Param("barcode")
	if productID == "" {
		h.logger.Errorf("[GetPdf] Couldn't get product id from query param")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get product id from query param",
		})
		return
	}

	ctxID, ok := c.Get(userIDCtx)
	if !ok {
		h.logger.Errorf("[GetPdf] Couldn't get user id from context")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get your credentials, try later",
		})
		return
	}

	userID, ok := ctxID.(int)
	if !ok {
		h.logger.Errorf("[GetPdf] Couldn't parse user id to int")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get your credentials, try later",
		})
		return
	}

	product, err := h.service.RetrieveProduct(productID, userID)
	if err != nil {
		h.logger.Errorf("[GetPdf] Error while retrieving product: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	loadedPDF, err := h.pdfService.LoadPDF(userID, product.Barcode)
	if err != nil {
		h.logger.Warnf("[GetPdf] Error while load pdf attempt: %s", err.Error())

		generatedPDF, err := h.pdfService.GeneratePDF(userID, *product)
		if err != nil {
			h.logger.Errorf("[GetPdf] Error while generating pdf: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.File(generatedPDF)
		return
	}

	c.File(loadedPDF)
}
