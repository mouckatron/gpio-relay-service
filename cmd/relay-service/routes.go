// Where to put all your service specific stuff!
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addRoutes() {
	ginRouter.GET("/relay", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		fmt.Printf("%+v", relays)

		_relays := make([]relay, 0)
		for _, r := range relays {
			_relays = append(_relays, *r)
		}

		c.JSON(http.StatusOK, _relays)
	})

	ginRouter.POST("/relay", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		gpio, _ := strconv.Atoi(c.PostForm("gpio"))
		name := c.PostForm("name")

		relayID := addRelay(gpio, name)

		c.JSON(http.StatusOK, relayID)
	})

	ginRouter.GET("/relay/:id", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		if r, ok := relays[c.Param("id")]; ok {
			c.JSON(http.StatusOK, r)
		} else {
			c.JSON(http.StatusNotFound, "")
		}
	})

	ginRouter.DELETE("/relay/:id", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		if _, ok := relays[c.Param("id")]; ok {
			delete(relays, c.Param("id"))
			c.JSON(http.StatusOK, "")
		} else {
			c.JSON(http.StatusNotFound, "")
		}
	})

	ginRouter.GET("/relay/:id/state", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		if r, ok := relays[c.Param("id")]; ok {
			c.JSON(http.StatusOK, r.State)
		} else {
			c.JSON(http.StatusNotFound, "")
		}
	})

	ginRouter.PUT("/relay/:id/state", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		if r, ok := relays[c.Param("id")]; ok {

			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "")
				return
			}

			if state, err := strconv.Atoi(string(body)); err == nil {
				if r.setState(state) {
					c.JSON(http.StatusOK, r.State)
				} else {
					c.JSON(http.StatusInternalServerError, "")
				}
			} else {
				c.Header("X-Error", "Incorrect state value: should be 0 or 1")
				c.JSON(http.StatusNotFound, "")
			}

		} else {
			c.JSON(http.StatusNotFound, "")
		}
	})
}
