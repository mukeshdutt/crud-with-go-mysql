package main

type EmployeeModel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type ResponseModel struct {
	IsSuccess   bool        `json:"isSuccess"`
	Message     string      `json:"message,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	SystemError string      `json:"systemError,omitempty"`
}
