package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DayNotes struct {
	Summary []string `json:"summary"`
	Details string   `json:"details"`
}

var calendarNotes = make(map[int]DayNotes)

func getDayNotes(c *gin.Context) {

	year, err := strconv.Atoi(c.DefaultQuery("year", "2025"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Query param for year is missing or not a number",
		})
		return
	}

	month, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Query param for month is missing or not a number",
		})
		return
	}

	fmt.Println(year, month)

	// Mock data, which should instead be collected from db
	calendarNotes[3] = DayNotes{
		Summary: []string{"go"},
		Details: "A Tour of Go: variables & functions",
	}
	calendarNotes[4] = DayNotes{
		Summary: []string{"go"},
		Details: "A Tour of Go: flow control - for",
	}
	calendarNotes[5] = DayNotes{
		Summary: []string{"go", "python"},
		Details: "A Tour of Go: flow control - if, else, switch, defer\n" +
			"python - enumerate\n" + "c# closures\n" + "java enhanced for loop",
	}
	calendarNotes[6] = DayNotes{
		Summary: []string{"go"},
		Details: "A Tour of Go: structs",
	}
	calendarNotes[10] = DayNotes{
		Summary: []string{"go"},
		Details: "A Tour of Go: slices and maps",
	}
	calendarNotes[11] = DayNotes{
		Summary: []string{"go"},
		Details: "A Tour of Go: exercise",
	}
	calendarNotes[12] = DayNotes{
		Summary: []string{"go"},
	}
	calendarNotes[17] = DayNotes{
		Summary: []string{"go"},
	}
	calendarNotes[18] = DayNotes{
		Summary: []string{"go"},
		Details: "A Tour of Go: ",
	}
	calendarNotes[20] = DayNotes{
		Summary: []string{"js", "css", "html"},
		Details: "creating a calendar with js",
	}
	calendarNotes[25] = DayNotes{
		Summary: []string{"css"},
		Details: "css universal selector and combinators",
	}
	calendarNotes[26] = DayNotes{
		Summary: []string{"js", "html"},
		Details: "figuring out how to display notes for the note calendar project",
	}

	c.IndentedJSON(http.StatusOK, calendarNotes)
}

func main() {
	router := gin.Default()
	router.GET("/daynotes", getDayNotes)

	router.Run("localhost:8080")
}
