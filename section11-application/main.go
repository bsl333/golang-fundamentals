package main

import (
	"course-material/section11-application/jsonDb"
	"course-material/section11-application/marshaling"
)

func main() {
	marshaling.Marshaling()
	jsonDb.RetrieveUsers()
}
