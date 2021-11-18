package repositories

type Movie struct {
	RecId        int    `db:"rec_id" json:"rec_id" form:"rec_id"`
	MovieId      string `db:"mov_id" json:"mov_id" form:"mov_id"`
	MovieRelease string `db:"mov_release" json:"mov_release" form:"mov_release"`
	ImgCover     string `db:"img_cover" json:"img_cover" form:"img_cover"`
}

type MoviesRepo interface {
	AddMovie(movie Movie) (*Movie, error)
	MovieAll() (movies []Movie, err error)
}
