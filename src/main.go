package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// _ "github.com/joho/godotenv/autoload"
	// "github.com/bwmarrin/discordgo"
)

// satoruchan ver:discord

var (
	Token string
)

func main() {

	router := gin.Default()
	t := time.Now()
	fmt.Println("Start App :",t)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv("DISCORD_TOKEN")

	router.GET("/", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "satoruchan start",
		})
	})

	fmt.Println(env)
	router.Run(":3030")

	// TODO: hooks discord room member(active)
	router.POST("/checkMember", func (c *gin.Context)  {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		id.First()
	})

	// TODO: hooks discord room member(diconnect)

	// TODO: check member name

	// TODO: post name and message for slack room
}