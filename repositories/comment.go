package repositories

import (
	"database/sql"
	"errors"
	commentModel "final-project/models/comment"
)

func GetAllComment(db *sql.DB, comment commentModel.Comment) (err error, result []commentModel.Comment) {
	sqlStatement := "SELECT * FROM comments SET WHERE article_id = $1"

	rows, err := db.Query(sqlStatement, comment.Article_id)
	if err != nil {
		return errors.New("data not found"), result
	}

	defer rows.Close()

	for rows.Next() {
		var comments = commentModel.Comment{}

		err = rows.Scan(&comments.ID, &comments.User_id, &comments.Article_id, &comments.Description, &comments.Created_at)
		if err != nil {
			panic(err)
		}

		result = append(result, comments)
	}

	return
}

func GetCommentDetail(db *sql.DB, comment commentModel.Comment) (err error, result []commentModel.Comment) {
	sqlStatement := "SELECT * FROM comments SET WHERE id = $1"

	row := db.QueryRow(sqlStatement, comment.ID)

	err = row.Scan(&comment.ID, &comment.User_id, &comment.Article_id, &comment.Description, &comment.Created_at)
	switch err {
	case sql.ErrNoRows:
		return errors.New("data not found"), result
	case nil:
		result = append(result, comment)
	default:
		return errors.New("data not found"), result
	}
	return
}

func InsertComment(db *sql.DB, comment commentModel.Comment, userId float64) (err error) {
	sqlStatement := "INSERT INTO comments (user_id, article_id, description) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sqlStatement, userId, comment.Article_id, comment.Description)

	return errs.Err()
}

func DeleteComment(db *sql.DB, comment commentModel.Comment, userId float64) (error) {
	sql := "DELETE FROM comments WHERE id = $1 AND user_id = $2"

	err := db.QueryRow(sql, comment.ID, userId)

	return err.Err()
}