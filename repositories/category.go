package repositories

import (
	"database/sql"
	"errors"
	categoryModel "final-project/models/category"
	"time"
)

func GetAllCategory(db *sql.DB) (err error, result []categoryModel.Category) {
	sql := "SELECT * FROM categories"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	

	defer rows.Close()

	for rows.Next() {
		var category = categoryModel.Category{}

		err = rows.Scan(&category.ID, &category.Name, &category.Created_at, &category.Updated_at)
		if err != nil {
			panic(err)
		}

		result = append(result, category)
	}
	return
}

func GetCategoryDetail(db *sql.DB, category categoryModel.Category) (err error, result []categoryModel.Category) {
	sqlStatement := "SELECT * FROM categories SET WHERE id = $1"

	row := db.QueryRow(sqlStatement, category.ID)

	err = row.Scan(&category.ID, &category.Name, &category.Created_at, &category.Updated_at)
	switch err {
	case sql.ErrNoRows:
		return errors.New("data not found"), result
	case nil:
		result = append(result, category)
	default:
		return errors.New("data not found"), result
	}
	return
}

func InsertCategory(db *sql.DB, category categoryModel.Category) (err error) {
	sql := "INSERT INTO categories (name) VALUES ($1)"

	errs := db.QueryRow(sql, category.Name)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category categoryModel.Category) (err error) {
	sql := "UPDATE categories SET name = $1, updated_at = $2 WHERE id = $3"

	updated_at := time.Now()
	errs := db.QueryRow(sql, category.Name, updated_at, category.ID)
	
	return errs.Err()
}

func DeleteCategory(db *sql.DB, category categoryModel.Category) (error) {
		// sql := "DELETE FROM articles WHERE category_id = $1"

		// err := db.QueryRow(sql, category.ID)
		// if err != nil {
		// 	return err.Err()
		// }

		sql := "DELETE FROM categories WHERE id = $1"
		err := db.QueryRow(sql, category.ID)

		return err.Err()
}
