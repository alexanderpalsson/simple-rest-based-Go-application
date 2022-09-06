package main

import (
	"errors"
	"fmt"
	"github.com/counter/internal/counters"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const counterName = "name"

var countersSingleton CountersStore

type CountersStore interface {
	Create(name string) error
	Increment(name string) error
	GetOne(name string) (int, error)
	GetAll() map[string]int
}

func main() {
	countersSingleton = counters.New()

	r := gin.Default() // Initiates gin server at port :8080

	r.POST(fmt.Sprintf("/create/:%s", counterName), createHandler)
	r.PUT(fmt.Sprintf("/increase/:%s", counterName), increaseHandler)
	r.GET("/get", getAllHandler)
	r.GET(fmt.Sprintf("/get/:%s", counterName), getOneHandler)

	if err := r.Run(); err != nil {
		fmt.Println("error running http handler", err)
	}
}

// TODO move routes to separate folder ./server/routes

func createHandler(c *gin.Context) {
	name, ok := parseNameParam(c)
	if !ok {
		return
	}

	if err := countersSingleton.Create(name); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusCreated)
}

func increaseHandler(c *gin.Context) {
	name, ok := parseNameParam(c)
	if !ok {
		return
	}

	err := countersSingleton.Increment(name)
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func getOneHandler(c *gin.Context) {
	name, ok := parseNameParam(c)
	if !ok {
		return
	}

	counter, err := countersSingleton.GetOne(name)
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{name: counter})
}

func getAllHandler(c *gin.Context) {
	c.JSON(http.StatusOK, countersSingleton.GetAll())
}

func parseNameParam(c *gin.Context) (name string, ok bool) {
	name = c.Param(counterName)

	if len(name) == 0 {
		handleError(c, http.StatusBadRequest, errors.New("missing counter name"))
		return "", false
	}

	return name, true
}

func handleError(c *gin.Context, statusCode int, err error) {
	log.Println(fmt.Errorf("error handling rest call %w", c.AbortWithError(statusCode, err)))
}
