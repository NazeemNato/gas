package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nazeemnato/gradient-avatar/src/utils"
)

func AvatarController(c *fiber.Ctx) error {
	// get the query params
	// getting name query from n
	n := c.Query("n")
	// getting size query from s
	s := c.Query("s")
	// checking if name is empty
	if len(n) == 0 {
		// if empty return 400 bad request error
		// because in order to create avatar we need name
		return c.Status(400).JSON(fiber.Map{
			"error": "name is required",
		})
	}
	// checking if size is empty
	// size is not required at all
	if len(s) == 0 {
		// if empty we set size to default 200px width and height
		s = "200"
	}
	if _, err := strconv.ParseInt(s,10,64); err  != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "size must be a number",
		})
	}
	// checking if size is less than 100px
	if sInt, _ := strconv.ParseInt(s,10,64); sInt >= 2000 {
		return c.Status(400).JSON(fiber.Map{
			"error": "size must be lesstan than 2000px",
		})
	}
	// converting name to hash integer
	hash := utils.Hash(n)
	// creating avatar from hash and size
	svg := utils.Avatar(hash, s)
	// set response type to svg
	c.Type("svg")
	// render svg
	return c.Send([]byte(svg))
}
