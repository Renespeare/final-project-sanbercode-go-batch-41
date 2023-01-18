package repositories

import (
	"database/sql"
	"errors"
	articleModel "final-project/models/article"
	"time"
)

func GetAllArticle(db *sql.DB) (err error, result []articleModel.Article) {
	sql := "SELECT * FROM articles"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var article = articleModel.Article{}

		err = rows.Scan(&article.ID, &article.User_id, &article.Category_id, &article.Title, &article.Description, &article.Created_at, &article.Updated_at)
		if err != nil {
			panic(err)
		}

		result = append(result, article)
	}
	return 
}

func GetArticleDetail(db *sql.DB, article articleModel.Article) (err error, result []articleModel.Article) {
	sqlStatement := "SELECT * FROM articles SET WHERE id = $1"

	row := db.QueryRow(sqlStatement, article.ID)

	err = row.Scan(&article.ID, &article.User_id, &article.Category_id ,&article.Title, &article.Description, &article.Created_at, &article.Updated_at)
	switch err {
	case sql.ErrNoRows:
		return errors.New("data not found"), result
	case nil:
		result = append(result, article)
	default:
		return errors.New("data not found"), result
	}
	return
}

func InsertArticle(db *sql.DB, article articleModel.Article, userId float64) (err error) {
	sql := "INSERT INTO articles (user_id, category_id, title, description) VALUES ($1, $2, $3, $4)"

	errs := db.QueryRow(sql, userId,  article.Category_id, article.Title, article.Description)

	return errs.Err()
}

func UpdateArticle(db *sql.DB, article articleModel.Article, userId float64) (err error) {
	sql := "UPDATE articles SET user_id = $1, category_id = $2, title = $3, description = $4, updated_at = $5 WHERE id = $6"

	updatedAt := time.Now()
	errs := db.QueryRow(sql, userId, article.Category_id ,article.Title, article.Description, updatedAt, article.ID)
	
	return errs.Err()
}

func DeleteArticle(db *sql.DB, article articleModel.Article, userId float64) (error) {
	// sql := "DELETE FROM comments WHERE article_id = $1"
	// err := db.QueryRow(sql, article.ID)
	// if err != nil {
	// 	return err.Err()
	// }

	sql := "DELETE FROM articles WHERE id = $1 AND user_id = $2"

	err := db.QueryRow(sql, article.ID, userId)

	return err.Err()
}