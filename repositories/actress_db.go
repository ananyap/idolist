package repositories

import (
	"github.com/jmoiron/sqlx"
)

type actressDb struct {
	db *sqlx.DB
}

func NewActDb(db *sqlx.DB) ActressRepo {
	return actressDb{db}
}

func (actDb actressDb) AddAct(act Actress) (*Actress, error) {

	query := "INSERT actress (act_name, act_jp_name, birth_date, tall, boob_cup, waist, hip, display) VALUES (?,?,?,?,?,?,?,?)"
	result, err := actDb.db.Exec(query, act.ActName, act.ActJpName, act.BirthDate, act.Tall, act.BoobCup, act.Waist, act.Hip, act.Display)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	actRes := Actress{
		ActId:     int(id),
		ActName:   act.ActName,
		ActJpName: act.ActJpName,
		BirthDate: act.BirthDate,
		Tall:      act.Tall,
		BoobCup:   act.BoobCup,
		Waist:     act.Waist,
		Hip:       act.Hip,
		Display:   act.Display,
	}

	return &actRes, nil
}

func (actDb actressDb) UpdateAct(actId int, actRequest Actress) (*Actress, error) {

	query := "UPDATE actress SET act_name=?, act_jp_name=?, birth_date=?, tall=?, boob_cup=?, waist=?, hip=?, display=? WHERE act_id=?"
	result, err := actDb.db.Exec(query, actRequest.ActName, actRequest.ActJpName, actRequest.BirthDate, actRequest.Tall, actRequest.BoobCup, actRequest.Waist, actRequest.Hip, actRequest.Display, actId)
	if err != nil {
		return nil, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected <= 0 {
		return nil, err
	}

	actRes := Actress{
		ActId:     actId,
		ActName:   actRequest.ActName,
		ActJpName: actRequest.ActJpName,
		BirthDate: actRequest.BirthDate,
		Tall:      actRequest.Tall,
		BoobCup:   actRequest.BoobCup,
		Waist:     actRequest.Waist,
		Hip:       actRequest.Hip,
		Display:   actRequest.Display,
	}
	return &actRes, nil
}
func (repo actressDb) DeleteAct(actRequest Actress) (*Actress, error) {

	query := "DELETE FROM actress WHERE act_id=?"
	res, err := repo.db.Exec(query, actRequest.ActId)
	if err == nil {
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, err
		}
		if affected <= 0 {
			return nil, err
		}

	}

	return &actRequest, nil
}
func (actDb actressDb) ActAll() ([]Actress, error) {
	actRes := []Actress{}
	query := "SELECT * FROM actress ORDER BY act_id ASC"
	err := actDb.db.Select(&actRes, query)
	if err != nil {
		return nil, err
	}

	return actRes, nil
}
func (actDb actressDb) ActById(atcId int) (*Actress, error) {
	act := Actress{}
	query := "SELECT * FROM actress WHERE act_id=?"
	err := actDb.db.Get(&act, query, atcId)
	if err != nil {
		return nil, err
	}

	return &act, err
}
