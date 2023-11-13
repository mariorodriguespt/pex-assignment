package main

import (
	"awesomeProject/database"
	"fmt"
)

type Visitor struct {
	id             string
	previousNumber int64
	currentNumber  int64
}

func createVisitor(newVisitor Visitor) error {
	sqlStatement := `
		INSERT INTO visitors (uuid, previous, current)
		values ($1, $2, $3);
`
	_, err := database.DB.Exec(sqlStatement, newVisitor.id, newVisitor.previousNumber, newVisitor.currentNumber)

	return err
}

func updateVisitor(newVisitor Visitor) error {
	sqlStatement := `
		UPDATE visitors
		SET previous=$1, current=$2
		WHERE uuid = $3;	
	`

	_, err := database.DB.Exec(sqlStatement, newVisitor.previousNumber, newVisitor.currentNumber, newVisitor.id)

	return err
}

func getVisitorById(visitorId any) *Visitor {
	var visitor Visitor

	fmt.Println("going for:", visitorId)
	rows, err := database.DB.Query(`SELECT * FROM visitors WHERE uuid = ?`, visitorId)

	if err != nil {
		fmt.Println("Failed to query", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&visitor.id, &visitor.previousNumber, &visitor.currentNumber)
		if err != nil {
			fmt.Println("Failed to get visitor ", err)
		}
		fmt.Println(visitor)
		return &visitor
	}

	return nil
}
