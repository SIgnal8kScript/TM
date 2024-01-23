package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Why Are You Here ? if you get here contact me you will work with me")
	})
	r.GET("/fasel", GETFaselUrl)
	r.Run(":8080")
}

func GETFaselUrl(c *gin.Context) {
	id := c.Query("id")
	playback := c.Query("playback")

	url := "https://www.faselplus.com/api/stream?&id=" + id + "&playback=" + playback
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("cookie", "_faselplus-access-token=G8Ebr3hUEgANhnOq7PPm8CgLALXmrbIkrhIPn5FPWPA=;_faselplus-profile-id=90ee1a10-4d09-d375-1cbb-2bad84ed4fa2")
	req.Header.Add("User-Agent", "okhttp/4.9.3")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	c.String(http.StatusOK, string(body))
}
