package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Filename string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {

}

// controllers
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"hello": "123",
		"test":  "aman",
	}

	err := json.NewEncoder(w).Encode(response) // response send as json value
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func getOneCours(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	// get params from url
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r) // all variables
	fmt.Println("Params", params)

	userId := params["id"]
	fmt.Println("userID", userId)
	response := map[string]string{
		"hello": "123",
		"test":  "aman",
	}

	err := json.NewEncoder(w).Encode(response) // response send as json value
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course                           // of type struct
	_ = json.NewDecoder(r.Body).Decode(&course) // reference of what to decode
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// we have all values inside course

	// rand.Seed(time.Now().UnixName)
	course.CourseId = strconv.Itoa(rand.Intn(100)) // convert int to string

	// push to courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return // no use of it but for safe side
}

// Update Request
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// Get Id of course that need to be updated
	// grab id from request
	params := mux.Vars(r) // all variables
	fmt.Println("Params", params)

	// Now loop through all values in courses slice and update its value
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // removing one element at the given index
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Update success")
			return
		}
	}

	json.NewEncoder(w).Encode("No course found")
	return

}
