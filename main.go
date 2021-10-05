package main

// https://levelup.gitconnected.com/consuming-a-rest-api-using-golang-b323602ba9d8

//POST
// curl http://localhost:8080/students `
// --include `
// --header "Content-Type: application/json" `
// --request "POST" `
// --data '{"StudentID": 7,"StudentName": "Ahmed","Enrollment": "new","CourseWorkLoad": 10.0}'


//go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
//go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type student struct {
	StudentID      int     `json:"studentid"`
	StudentName    string  `json:"studentname"`
	Enrollment     string  `json:"enrollment"`
	CourseWorkLoad float64 `json:"courseworkload"`
}

type course struct {
	CourseID      int     `json:"courseid"`
	CourseName    string  `json:"coursename"`
	CourseTeacher teacher `json:"courseteacher"`
	CourseScore   float64 `json:"coursescore"`
}

type teacher struct {
	TeacherID     int     `json:"teacherid"`
	TeacherName   string  `json:"teachername"`
	TeacherRating float64 `json:"teacherrating"`
}

func main() {
	fmt.Println("Calling API...")
	fmt.Println("")

	inputSwitch()

}

func inputSwitch() {
	fmt.Println("Write your request:\nGET, PUT, DELETE or POST")
	fmt.Print("Input: ")
	var input string
	fmt.Scanln(&input)

	switch input {
	case "GET":
		getAllCourses()
	case "PUT":
		fmt.Println("PUT not implemented")
	case "DELETE":
		fmt.Println("DELETE not implemented")
	case "POST":
		fmt.Println("POST not implemented")
	}

}

func getAllCourses() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/courses", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var courseObject []course
	json.Unmarshal(bodyBytes, &courseObject)
	fmt.Printf("%+v\n", courseObject)
}
