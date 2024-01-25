package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Seat represents the Seat table structure
type Seat struct {
	ID      int
	Student string
}

func main() {
	// Assuming you have a database connection available
	db, err := sql.Open("mysql", "root:Sanju@123@tcp(localhost:3306)/students")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Run the swapSeats function
	err = swapSeats(db)
	if err != nil {
		log.Fatal(err)
	}
}

func swapSeats(db *sql.DB) error {
	// Fetch all rows from the Seat table
	rows, err := db.Query("SELECT id, student FROM Seat ORDER BY id")
	if err != nil {
		return err
	}
	defer rows.Close()

	var seats []Seat

	// Iterate through the rows and store the data in the seats slice
	for rows.Next() {
		var seat Seat
		err := rows.Scan(&seat.ID, &seat.Student)
		if err != nil {
			return err
		}
		seats = append(seats, seat)
	}

	// Swap the seat IDs for consecutive students (excluding the last one if the number is odd)
	for i := 0; i < len(seats)-1; i += 2 {
		seats[i].ID, seats[i+1].ID = seats[i+1].ID, seats[i].ID
	}

	// Update the Seat table with the swapped IDs
	for _, seat := range seats {
		_, err := db.Exec("UPDATE Seat SET id = ? WHERE student = ?", seat.ID, seat.Student)
		if err != nil {
			return err
		}
	}

	// Fetch and print the updated Seat table
	rows, err = db.Query("SELECT id, student FROM Seat ORDER BY id")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Output:")
	fmt.Println("+----+---------+")
	fmt.Println("| id | student |")
	fmt.Println("+----+---------+")
	for rows.Next() {
		var seat Seat
		err := rows.Scan(&seat.ID, &seat.Student)
		if err != nil {
			return err
		}
		fmt.Printf("| %2d | %-7s |\n", seat.ID, seat.Student)
	}
	fmt.Println("+----+---------+")

	return nil
}
