r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/profile", func(c *gin.Context) {
		playerName := c.Query("name")

		var scores []int
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
	})

	r.POST("/new_score", func(c *gin.Context) {
		var playthrough record
		if c.BindJSON(&playthrough) != nil {
			c.Status(400)
			return
		}

		db.Exec("INSERT INTO mtmScores (playerName, scores) VALUES ($1, $2)", playthrough.Name, playthrough.Score)

		c.String(201, "Created record by "+playthrough.Name+" with score "+fmt.Sprint(playthrough.Score))
		//writeDB()
	})

	r.GET("/score_table", func(c *gin.Context) {
		var playthrough record
		var users []player

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
				var user player = player{playthrough.Name, []int{playthrough.Score}}
				users = append(users, user)
			}
		}
		err = rows.Err()

		for i := range users {
			sort.Sort(sort.Reverse(sort.IntSlice(users[i].Score)))
		}

		c.JSON(200, users)
	})
