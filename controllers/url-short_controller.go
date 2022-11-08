package controllers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/url"
	"time"
	"url-shortener-back/configs"
	"url-shortener-back/models"
	"url-shortener-back/responses"
	"url-shortener-back/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var urlCollection *mongo.Collection = configs.GetCollection(configs.DB, "urls-short")
var validate = validator.New()

func CreateURLShort() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var urlShort models.URLShort
		defer cancel()

		if err := c.BindJSON(&urlShort); err != nil {
			c.JSON(http.StatusBadRequest, responses.URLShortResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&urlShort); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.URLShortResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		if _, validationErr := url.ParseRequestURI(urlShort.Url); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.URLShortResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		if urlShort.Id == "AUTO" {
			urlShort.Id = utils.RandomString(8)
		}else {
			urlShort.Id = utils.ValidIDString(urlShort.Id)
		}

		//TODO: Validar si el id enviado o generado no existe en la BD.
		newURLShort := models.URLShort{
			Id:      urlShort.Id,
			Url:     urlShort.Url,
		}

		result, err := urlCollection.InsertOne(ctx, newURLShort)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.URLShortResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var resultMap = make(map[string]string)
		resultMap["InsertedID"] = fmt.Sprint(result.InsertedID)
		resultMap["UrlShortID"] = urlShort.Id
		c.JSON(http.StatusCreated, responses.URLShortResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"result": resultMap}})
	}
}


func GetURLFull() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		URLShortId := c.Param("urlId")
		var urlShort models.URLShort
		defer cancel()

		err := urlCollection.FindOne(ctx, bson.M{"id": URLShortId}).Decode(&urlShort)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				c.JSON(http.StatusNotFound, responses.URLShortResponse{Status: http.StatusNotFound, Message: "not found", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusInternalServerError, responses.URLShortResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.URLShortResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": urlShort}})
	}
}

func GetAllURLs() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var urlShort []models.URLShort
		defer cancel()

		results, err := urlCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.URLShortResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.URLShort
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.URLShortResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			urlShort = append(urlShort, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.URLShortResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": urlShort}},
		)
	}
}