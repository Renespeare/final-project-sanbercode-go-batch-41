package main

import (
	"database/sql"
	"final-project/controllers"
	"final-project/database"
	"final-project/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
	err error
)

func main() {

	r := gin.Default()

	// env config
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed load file environment")
	} else {
		fmt.Println("Success read file environment")
	}
	
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))
	// fmt.Println(psqlInfo)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()


	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
  public.Use(middleware.Authentication())
	{
		// logout
		public.POST("/logout", controllers.Logout)

		// category
		public.GET("/category", controllers.GetAllCategory)
		public.GET("/category/:id", controllers.GetCategoryDetail)
		public.POST("/category", controllers.InsertCategory)
		public.PUT("/category/:id", controllers.UpdateCategory)
		public.DELETE("/category/:id", controllers.DeleteCategory)

		// article
		public.GET("/article", controllers.GetAllArticle)
		public.GET("/article/:id", controllers.GetArticleDetail)
		public.POST("/article", controllers.InsertArticle)
		public.PUT("/article/:id", controllers.UpdateArticle)
		public.DELETE("/article/:id", controllers.DeleteArticle)

		// comment
		public.GET("/comment", controllers.GetAllComment)
		public.GET("/comment/:id", controllers.GetCommentDetail)
		public.POST("/comment", controllers.InsertComment)
		public.DELETE("/comment/:id", controllers.DeleteComment)

	}

	r.Run(":" + os.Getenv("PORT"))
	// r.Run(":8080")
}