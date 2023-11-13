package main

import (
	"awesomeProject/database"
	"awesomeProject/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func setValuesForVisitor(context *gin.Context, visitor Visitor) {
	err := updateVisitor(visitor)

	if err != nil {
		fmt.Println("error ", err)
		context.JSON(500, gin.H{
			"error": "operation failed",
		})
		return
	}

	context.JSON(200, gin.H{
		"number": visitor.currentNumber,
	})
}

func main() {
	err := database.InitDatabase()

	if err != nil {
		return
	}
	server := gin.Default()

	server.Use(middlewares.CORSMiddleware())
	publicRoutes := server.Group("/")
	publicRoutes.Use(middlewares.CheckIfVisitorAlreadyExists)
	publicRoutes.GET("/current", func(context *gin.Context) {
		visitorId, _ := context.Get("visitor-id")

		if visitorId == nil {
			newVisitorId := uuid.New().String()

			var err = createVisitor(Visitor{
				id:             newVisitorId,
				previousNumber: 0,
				currentNumber:  0,
			})
			if err != nil {
				fmt.Println("Failed to create visitor ", err)
				return
			}

			context.JSON(200, gin.H{
				"visitorId": newVisitorId,
				"number":    0,
			})

		} else {
			visitor := getVisitorById(visitorId)

			context.JSON(200, gin.H{
				"number": visitor.currentNumber,
			})
		}
	})

	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middlewares.IdentifyUser)
	authenticatedRoutes.POST("/next", func(context *gin.Context) {
		fmt.Println("Next endpoint getting called???")
		visitorId, _ := context.Get("visitor-id")
		visitor := getVisitorById(visitorId)

		if visitor == nil {
			context.JSON(500, gin.H{
				"error": "Unknown visitor",
			})
			return
		}

		if visitor.previousNumber == 0 && visitor.currentNumber == 0 {
			setValuesForVisitor(context, Visitor{
				id:             visitor.id,
				previousNumber: 0,
				currentNumber:  1,
			})
			return
		} else if visitor.previousNumber == 0 && visitor.currentNumber == 1 {
			setValuesForVisitor(context, Visitor{
				id:             visitor.id,
				previousNumber: 1,
				currentNumber:  1,
			})
			return
		}

		setValuesForVisitor(context, Visitor{
			id:             visitor.id,
			previousNumber: visitor.currentNumber,
			currentNumber:  visitor.currentNumber + visitor.previousNumber,
		})
	})

	authenticatedRoutes.POST("/previous", func(context *gin.Context) {
		visitorId, _ := context.Get("visitor-id")
		visitor := getVisitorById(visitorId)

		if visitor.previousNumber == 0 && visitor.currentNumber == 0 {
			context.JSON(200, gin.H{
				"number": 0,
			})
			return
		} else if visitor.previousNumber == 0 && visitor.currentNumber == 1 {
			setValuesForVisitor(context, Visitor{
				id:             visitor.id,
				previousNumber: 0,
				currentNumber:  0,
			})
			return
		}

		setValuesForVisitor(context, Visitor{
			id:             visitor.id,
			previousNumber: visitor.currentNumber - visitor.previousNumber,
			currentNumber:  visitor.previousNumber,
		})
	})

	err = server.Run()
	if err != nil {
		fmt.Println("Failed to launch web server")
		return
	}
}
