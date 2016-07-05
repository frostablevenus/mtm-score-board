package logicHandler

import (
	"database/sql"

	"mtm-score-board/resources/models/response"
)

func GetPlayThrough(db *sql.DB, playerName string) (*sql.Rows, error) {
	return db.Query("SELECT scores FROM mtmScores WHERE playerName=$1", playerName)
}

func CreatePlayThrough(db *sql.DB, playthrough *response.Record) {
	db.Exec("INSERT INTO mtmScores (playerName, scores) VALUES ($1, $2)", playthrough.Name, playthrough.Score)
}

func ListPlayThrough(db *sql.DB) (*sql.Rows, error) {
	return db.Query("SELECT playerName, scores FROM mtmScores")
}
