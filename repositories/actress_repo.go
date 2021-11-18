package repositories

type Actress struct {
	ActId     int    `db:"act_id" json:"act_id"`
	ActName   string `db:"act_name" json:"act_name"`
	ActJpName string `db:"act_jp_name" json:"act_jp_name"`
	BirthDate string `db:"birth_date" json:"birth_date"`
	Tall      int    `db:"tall" json:"tall"`
	BoobCup   string `db:"boob_cup" json:"boob_cup"`
	Waist     int    `db:"waist" json:"waist"`
	Hip       int    `db:"hip" json:"hip"`
	Display   string `db:"display" json:"display"`
}

type ActressRepo interface {
	AddAct(act Actress) (*Actress, error)
	UpdateAct(act int, actRequest Actress) (*Actress, error)
	DeleteAct(act Actress) (*Actress, error)
	ActAll() ([]Actress, error)
	ActById(atcId int) (*Actress, error)
}
