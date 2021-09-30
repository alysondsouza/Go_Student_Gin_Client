package main

// https://levelup.gitconnected.com/consuming-a-rest-api-using-golang-b323602ba9d8

//POST
//curl http://localhost:8080/students `
//--include `
//--header "Content-Type: application/json" `
//--request "POST" `
//--data '{"StudentID": 7,"StudentName": "Ahmed","Enrollment": "new","CourseWorkLoad": 10.0}'

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type student struct {
	StudentID      int     `json:"studentid"` //STRING VS INT
	StudentName    string  `json:"studentname"`
	Enrollment     string  `json:"enrollment"` //STRING VS ENUM
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
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
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

	var studentObject student
	json.Unmarshal(bodyBytes, &studentObject)
	fmt.Printf("API Response as struct %+v\n", studentObject)
}
