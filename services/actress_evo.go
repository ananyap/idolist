package services

import (
	"fmt"

	"github.com/ananyap/idolist/repositories"
	"github.com/ananyap/idolist/utils"
	"github.com/gofiber/fiber/v2"
)

type actressEvo struct {
	actRepo repositories.ActressRepo
}

func NewActressEvo(repo repositories.ActressRepo) ActressService {
	return actressEvo{repo}
}

func (repo actressEvo) GetActAll() ([]ActressEvo, error) {
	acts, err := repo.actRepo.ActAll()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity)
	}

	fmt.Println(acts)

	actsEvo := []ActressEvo{}
	for _, act := range acts {
		age, _ := utils.AgeFromDateOfBirth(act.BirthDate)
		actEvo := ActressEvo{
			ActName: act.ActName,
			Age:     age,
			Cup:     act.BoobCup,
			Tall:    act.Tall,
			Waist:   act.Waist,
			Hip:     act.Hip,
			Display: "https://live.staticflickr.com/65535/" + act.Display,
		}
		actsEvo = append(actsEvo, actEvo)
	}

	return actsEvo, nil
}

func (repo actressEvo) GetActEvoById(actId int) (*ActressEvo, error) {
	act, err := repo.actRepo.ActById(actId)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity)
	}

	age, _ := utils.AgeFromDateOfBirth(act.BirthDate)

	actEvo := ActressEvo{
		ActName: act.ActName,
		Age:     age,
		Cup:     act.BoobCup,
		Tall:    act.Tall,
		Waist:   act.Waist,
		Hip:     act.Hip,
		Display: "https://live.staticflickr.com/65535/" + act.Display,
	}
	return &actEvo, nil
}
