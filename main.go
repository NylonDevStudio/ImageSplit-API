package main

import (
	"image/png"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	gim "github.com/ozankasikci/go-image-merge"
)

func main() {
	router := gin.Default()
	router.GET("/tile/:w/:h/:id", makeImage)
	router.Run(":8080")
}

func makeImage(c *gin.Context) {
	id := c.Param("id")
	widthParam := c.Param("w")
	heightParam := c.Param("h")

	w, err := strconv.Atoi(widthParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid width parameter"})
		return
	}
	h, err := strconv.Atoi(heightParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid height parameter"})
		return
	}

	var images = strings.Split(id, "-")
	grids := make([]*gim.Grid, len(images))

	for i, path := range images {
		grids[i] = &gim.Grid{ImageFilePath: "images/" + path + ".png"}
	}

	rgba, err := gim.New(grids, w, h).Merge()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error merging images"})
		return
	}

	c.Header("Content-Type", "image/png")
	err = png.Encode(c.Writer, rgba)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error encoding image"})
	}
}
