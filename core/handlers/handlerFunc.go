package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"sort"

	"mtmScoreBoard/resources"
	"mtmScoreBoard/resources/models/response"
)

type PlaythroughHandler interface {
	GetPlaythrough(*gin.Context, string)
	CreatePlaythrough(*gin.Context, *response.Record)
	ListPlaythrough(*gin.Context)
}

type Handler struct {
	R *resources.Resource
}

func NewPlaythroughHandler(r *resources.Resource) PlaythroughHandler {
	return &Handler{
		R: r,
	}
}

func (h *Handler) GetPlaythrough(c *gin.Context, playerName string) {
	var scores []int
	db := h.R.PostgreSql

	rows, err := db.Query("SELECT scores FROM mtmScores WHERE playerName=$1", playerName)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var score int
		err = rows.Scan(&score)
		scores = append(scores, score)
	}
	err = rows.Err()

	switch {
	case err == sql.ErrNoRows:
		c.String(404, "Player does not exist")
	case err != nil:
		log.Fatal(err)
	default:
		sort.Sort(sort.Reverse(sort.IntSlice(scores)))
		c.JSON(200, gin.H{
			"Name":   playerName,
			"Scores": scores,
		})
	}
}

func (h *Handler) CreatePlaythrough(c *gin.Context, playthrough *response.Record) {
	db := h.R.PostgreSql

	db.Exec("INSERT INTO mtmScores (playerName, scores) VALUES ($1, $2)", playthrough.Name, playthrough.Score)

	c.String(201, "Created record by "+playthrough.Name+" with score "+fmt.Sprint(playthrough.Score))
}

func (h *Handler) ListPlaythrough(c *gin.Context) {
	var playthrough response.Record
	var users []response.Player

	fmt.Println(h.R == nil)
	fmt.Println(h.R.PostgreSql == nil)

	db := *h.R.PostgreSql

	rows, err := db.Query("SELECT playerName, scores FROM mtmScores")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&playthrough.Name, &playthrough.Score)

		temp := true
		for i := range users {
			if playthrough.Name == users[i].Name {

				users[i].Score = append(users[i].Score, playthrough.Score)
				temp = false
				break
			}
		}
		if temp == true {
			var user response.Player = response.Player{playthrough.Name, []int{playthrough.Score}}
			users = append(users, user)
		}
	}
	err = rows.Err()

	for i := range users {
		sort.Sort(sort.Reverse(sort.IntSlice(users[i].Score)))
	}

	c.JSON(200, users)

}
