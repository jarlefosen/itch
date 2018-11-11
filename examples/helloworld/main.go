package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
)

func main() {
	uid, username, err := currentUser()
	if err != nil {
		log.Fatalf("failed to get current user: %v", err)
	}
	gid, group, err := currentGroup()
	if err != nil {
		log.Fatalf("failed to get current group: %v", err)
	}

	fmt.Printf("Running as user (uid=%d) %s, group (gid=%d) %s\n", uid, username, gid, group)
	fmt.Printf("Hello World\n")
}

func currentUser() (int, string, error) {
	uid := os.Getuid()
	if uid == 0 {
		return uid, "root", nil
	}
	usr, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		return uid, "", err
	}
	return uid, usr.Name, nil
}

func currentGroup() (int, string, error) {
	gid := os.Getgid()
	if gid == 0 {
		return gid, "root", nil
	}
	group, err := user.LookupGroupId(strconv.Itoa(gid))
	if err != nil {
		return gid, "", err
	}
	return gid, group.Name, nil
}
