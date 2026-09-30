package main

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

func fetchUserData(userID int) error {
	userFound := true

	if !userFound {
		return &CustomError{
			Code:    404,
			Message: "user not found",
		}
	}

	return nil
}

func main() {
	err := fetchUserData(42)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("no error!")
	}
}
