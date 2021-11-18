package services

type ActressEvo struct {
	ActName string `json:"name"`
	Age     int    `json:"age"`
	Cup     string `json:"cup"`
	Tall    int    `json:"tall"`
	Waist   int    `json:"waist"`
	Hip     int    `json:"hip"`
	Display string `json:"display"`
}

type ActressService interface {
	GetActEvoById(actId int) (*ActressEvo, error)
	GetActAll() ([]ActressEvo, error)
}
