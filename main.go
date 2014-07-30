package main

import (
    "fmt"
    "strconv"
    // "weatherwars/request"
    "github.com/wsxiaoys/terminal/color"
)

func main() {

    var placeNames []string
    var placeName string

    color.Println("@M Note: @m Use an empty string to begin the war!")

    for i := 0; ; i++ {

        placeName = ""

        // Prompt the user to enter their place name.
        fmt.Printf("Place: ")
        fmt.Scanf("%s", &placeName)

        if (len(placeName) == 0) {

            // We've got all of the places we need to begin the process of comparing
            // the weather!
            break

        }

        // Voila! We have another place name so let's add it to the array of strings for later.
        placeNames = append(placeNames, placeName)

    }

    // At this point we're ready to begin the whole process.
    color.Println("@gPreparing the war with " + strconv.Itoa(len(placeNames)) + " places...")

}
