package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func getMonthNotesEndpoint(c *gin.Context) {

	year, err := strconv.Atoi(c.Param("year"))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Path param for year is missing or not a number",
		})
		return
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Path param for month is missing or not a number",
		})
		return
	}

	calendarNotes := getCalendarNotes(year, month)

	c.IndentedJSON(http.StatusOK, calendarNotes)
}

func putDayNotesEndpoint(c *gin.Context) {

	year, err := strconv.Atoi(c.Param("year"))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Path param for year is missing or not a number",
		})
		return
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Path param for month is missing or not a number",
		})
		return
	}

	day, err := strconv.Atoi(c.Param("day"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Path param for day is missing or not a number",
		})
		return
	}

	note := DayNotes{}

	if err := c.ShouldBindBodyWith(&note, binding.JSON); err == nil {

		// if note is empty, abort changes and inform the frontend to do a delete instead
		if (len(note.Summary) == 0 ||
			(len(note.Summary) == 1 && len(note.Summary[0]) == 0)) &&
			len(note.Details) == 0 {

			c.JSON(http.StatusNoContent, map[string]any{
				"success": false,
				"message": "Both note summary and details are empty. Do a DELETE request instead.",
			})
			return
		}

		updateCalendarNote(year, month, day, note)

		c.Status(http.StatusOK)

	} else {
		c.JSON(http.StatusNoContent, map[string]any{
			"success": false,
			"message": "Body (with JSON) for PUT request is missing or not of the form {summary: [...], details: ...}",
		})
		return
	}

}

func deleteDayNotesEndpoint(c *gin.Context) {

	year, err := strconv.Atoi(c.Param("year"))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Path param for year is missing or not a number",
		})
		return
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Path param for month is missing or not a number",
		})
		return
	}

	day, err := strconv.Atoi(c.Param("day"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Path param for day is missing or not a number",
		})
		return
	}

	deleteDayNotes(year, month, day)

	c.Status(http.StatusOK)
}

func main() {
	router := gin.Default()
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.StaticFile("/preview.html", "./preview.html")
	router.StaticFile("/notes-calendar/notes-calendar.js", "./notes-calendar/notes-calendar.js")
	router.StaticFile("/notes-calendar/notes-calendar.css", "./notes-calendar/notes-calendar.css")

	router.GET("/daynotes/:year/:month", getMonthNotesEndpoint)
	router.PUT("/daynotes/:year/:month/:day", putDayNotesEndpoint)
	router.DELETE("/daynotes/:year/:month/:day", deleteDayNotesEndpoint)

	populateDBsample()

	router.Run("localhost:8080")
}
