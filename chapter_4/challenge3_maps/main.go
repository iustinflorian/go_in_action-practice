package main

import "fmt"

func main() {
	permissions := map[string]string{
		"alice":   "admin",
		"bob":     "editor",
		"charlie": "",
	}

	checkAccess(permissions, "alice")

	delete(permissions, "alice")
	checkAccess(permissions, "alice")

	checkAccess(permissions, "charlies")
}

func checkAccess(permissions map[string]string, user string) {
	val, ok := permissions[user]
	if !ok {
		fmt.Printf("User %s not found\n", user)
		return
	}
	fmt.Printf("User %s with role %s found\n", user, val)
}
