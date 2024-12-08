package main

import (
	"fmt"
	"strconv"
	"errors"
)

const NUM_CHOICES = 4
func print_choices() {
	choices := [NUM_CHOICES]string{"Create entry", "Get all entries", "Update an entry", "Delete an entry"}

	for i := 0; i < len(choices); i++ {
		fmt.Printf("%d) %s\n", i + 1, choices[i])
	}
}

type ErrorCallback func()
// Reports an error if it exists, calls function cb
func report_if_error(err error, cb ErrorCallback) {
	if err != nil {
		fmt.Println("An error occured:", err)
		cb()
	}
}

func get_input() int {
	fmt.Println("Enter choice: ");

	var input string
	_, err := fmt.Scanln(&input)
	report_if_error(err, func(){})

	int_choice, err := strconv.Atoi(input)
	report_if_error(err, func(){})

	return int_choice
}

func is_valid_choice(choice int) bool {
	return (choice >= 1 && choice <= NUM_CHOICES)
}

func create_db() map[string]string {
	db := make(map[string]string)
	return db
}

func create_entry(db map[string]string, key string, val string) bool {
	db[key] = val
	if db[key] == "" {
		return false
	}
	return true
}

func get_entries(db map[string]string) {
	var i int = 1
	for key, value := range db {
		fmt.Printf("%d) %s %s\n", i, key, value)
		i++
	}
}

func update_entry(db map[string]string, key string, updated_val string) bool {
	db[key] = updated_val

	if db[key] != updated_val {
		return false
	}

	return true
}

func delete_entry(db map[string]string, key string) bool {
	if db[key] == "" {
		return false
	} else {
		delete(db, key)
		return true
	}
}

func program_loop() {
	db := create_db()

	print_choices()
	program_running := true
	for program_running {
		choice := get_input()
		for !is_valid_choice(choice) {
			report_if_error(errors.New("Invalid choice"), func(){})
			choice = get_input()
		}

		switch choice {
			case 1:
				var key, val string
				fmt.Println("Enter key:")
				fmt.Scanln(&key)
				fmt.Println("Enter value:")
				fmt.Scanln(&val)

				if !create_entry(db, key, val) {
					report_if_error(errors.New("Could not create entry"), func(){})
				} else {
					fmt.Println("Entry stored successfully!")
				}
			case 2:
				get_entries(db)
			case 3:
				var key, val string
				fmt.Println("Enter key:")
				fmt.Scanln(&key)
				fmt.Println("Enter value:")
				fmt.Scanln(&val)

				if !update_entry(db, key, val) {
					report_if_error(errors.New("Could not update entry"), func(){})
				} else {
					fmt.Println("Entry updated successfully!")
				}
			case 4:
				var key string
				fmt.Println("Enter key:")
				fmt.Scanln(&key)

				if !delete_entry(db, key) {
					report_if_error(errors.New("Could not delete entry"), func(){})
				} else {
					fmt.Println("Entry deleted successfully!")
				}
		}
	}
}

func main() {
	fmt.Println("Welcome to In-memory database written in Golang")
	program_loop()
}
