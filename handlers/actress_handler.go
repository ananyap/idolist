package handlers

import (
	"strconv"

	"github.com/ananyap/idolist/services"
	"github.com/gofiber/fiber/v2"
)

type ActressApi struct {
	Status string               `json:"status"`
	Data   *services.ActressEvo `json:"data"`
}

type ActressAllApi struct {
	Status string                `json:"status"`
	Data   []services.ActressEvo `json:"data"`
}

type actressHandler struct {
	actressService services.ActressService
}

func NewActressHandler(actService services.ActressService) actressHandler {
	return actressHandler{actService}
}

func (actHandler actressHandler) ActressAllHandler(c *fiber.Ctx) error {
	myMessage := ActressAllApi{}

	acts, err := actHandler.actressService.GetActAll()

	if err != nil {
		myMessage.Status = err.Error()
		myMessage.Data = nil
	} else {
		myMessage.Status = "OK"
		myMessage.Data = acts
	}

	return c.JSON(myMessage)
}

func (actHandler actressHandler) ActressHandler(c *fiber.Ctx) error {
	myMessage := ActressApi{}
	actId := c.Params("actid")
	i, _ := strconv.Atoi(actId)
	act, err := actHandler.actressService.GetActEvoById(i)
	if err != nil {
		myMessage.Status = err.Error()
		myMessage.Data = nil
	} else {
		myMessage.Status = "OK"
		myMessage.Data = act
	}

	return c.JSON(myMessage)
}
