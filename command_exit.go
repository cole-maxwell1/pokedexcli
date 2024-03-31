package main

import "os"

func callbackExit(state *State) error {
	os.Exit(0)
	return nil
}
