package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ananyap/idolist/handlers"
	"github.com/ananyap/idolist/repositories"
	"github.com/ananyap/idolist/services"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:IntelliP24@tcp(localhost:3306)/gravurelist")
	if err != nil {
		panic(err)
	}
	engine := html.NewFileSystem(http.Dir("./views"), ".html")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./views")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	app.Get("/moviesall/", func(c *fiber.Ctx) error {

		repo := repositories.NewMoviesDb(db)
		movs, err := repo.MovieAll()
		if err != nil {
			return err
		}

		return c.Render("page_mov_all", fiber.Map{

			"movies": movs,
			"alert":  "info",
		}, "layouts/main")
	})

	repo := repositories.NewActDb(db)
	serviec := services.NewActressEvo(repo)
	handler := handlers.NewActressHandler(serviec)

	app.Get("/actress/:actid", handler.ActressHandler)
	app.Get("/actressall", handler.ActressAllHandler)

	app.Get("/moviex/:start/:end", func(c *fiber.Ctx) error {
		startRec, _ := strconv.Atoi(c.Params("start"))
		endRec, _ := strconv.Atoi(c.Params("end"))

		repo := repositories.NewMoviesDb(db)
		movs, lastRec, pageNum, amount, err := repo.MovieAll(startRec, endRec)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"movies":  movs,
			"amount":  amount,
			"pageNum": pageNum,

			"lastRec": lastRec,
		})
	})

	app.Post("/movie", func(c *fiber.Ctx) error {
		fmt.Printf("IsJson: %v\n", c.Is("json"))
		fmt.Println(string(c.Body()))

		movieRequest := repositories.Movie{}

		err := c.BodyParser(&movieRequest)
		if err != nil {
			return err
		}

		fmt.Println(movieRequest)

		repoMovies := repositories.NewMoviesDb(db)
		movRes, err := repoMovies.AddMovie(movieRequest)

		if err != nil {
			return err
		}

		fmt.Println("movRes", movRes)

		//c.Redirect("/movies/0/20")

		return c.JSON(fiber.Map{
			"message": "upload complete",
			"alert":   "sucess",
		})

	})

	app.Get("/movie_form", func(c *fiber.Ctx) error {
		return c.Render("page_mov_form", fiber.Map{
			"alert": "info",
		}, "layouts/main")
	})

	app.Listen(":3000")

}
