package main

import (
    "github.com/gin-gonic/gin"
    "image/png"
    "log"
    "strings"
    "os"
    "strconv"

    gim "github.com/ozankasikci/go-image-merge"
)

type filePathStruct struct {
    ImageFilePath string
}

func main() {
    router := gin.Default()
    router.GET("/tile/:w/:h/:id", makeImage)
    
    router.Run("localhost:8088")
}

func makeImage(c *gin.Context) {
    id := c.Param("id");
    var s = id
    
    h, _ := strconv.Atoi("5")
    w, _ := strconv.Atoi("5")
    
    // create an array of the images we have
    var images = strings.Split(s, "-")

    // Create an array to store *gim.Grid instances
    grids := make([]*gim.Grid, len(images))

    // Iterate over the file paths and create *gim.Grid instances
    for i, path := range images {
        // we use a reference to the grid so we can modify the object outside the loop
        grids[i] = &gim.Grid{ImageFilePath: "images/" + path + ".png"}
    }
    
        // make the grid of images
    rgba, err := gim.New(grids, w, h).Merge()
    if err != nil {
        log.Fatal(err)
    }
    // create the image
    file, err := os.Create("./path.png")
    err = png.Encode(file, rgba)
}
