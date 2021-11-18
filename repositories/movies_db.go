package repositories

import (
	"github.com/jmoiron/sqlx"
)

type movieDb struct {
	db *sqlx.DB
}

func NewMoviesDb(db *sqlx.DB) MoviesRepo {
	return movieDb{db}
}

func (repo movieDb) AddMovie(movie Movie) (*Movie, error) {

	query := "INSERT movies (img_cover, mov_id, mov_release) VALUES (?,?,?)"
	result, err := repo.db.Exec(query, movie.ImgCover, movie.MovieId, movie.MovieRelease)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	movieRes := Movie{
		RecId:        int(id),
		MovieId:      movie.MovieId,
		MovieRelease: movie.MovieRelease,
		ImgCover:     movie.ImgCover,
	}

	return &movieRes, nil
}

func (repo movieDb) MovieAll() (movies []Movie, err error) {

	moviesRes := []Movie{}
	query := "SELECT * FROM movies ORDER BY rec_id DESC"
	err = repo.db.Select(&moviesRes, query)
	if err != nil {
		return nil, err
	}

	return moviesRes, nil
}
