package main

import (
	"encoding/json"
	"net/http"

	"k8s.io/klog/v2"
)

type Employee struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Payroll  int    `json:"payroll"`
	Gender   string `json:"gender"`
}

type Message struct {
	Employees   []Employee `json:"employee"`
	CompanyName string     `json:"company_name"`
	Address     string     `json:"address"`
	Phone       string     `json:"phone"`
}

type Response struct {
	Message Message `json:"message"`
	Status  string  `json:"status"`
}

func getEmployees() []Employee {
	john := Employee{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Birthday: "1993-01-01",
		Payroll:  50000,
		Gender:   "Male",
	}

	jane := Employee{
		Name:     "Jane Smith",
		Email:    "jane.smith@example.com",
		Birthday: "1995-05-15",
		Payroll:  60000,
		Gender:   "Female",
	}

	alice := Employee{
		Name:     "Alice Johnson",
		Email:    "alice.johnson@example.com",
		Birthday: "1988-03-10",
		Payroll:  70000,
		Gender:   "Female",
	}

	return []Employee{john, jane, alice}
}

func handler(w http.ResponseWriter, r *http.Request) {
	klog.Info("Received request, %w", r)

	response := Response{
		Message: Message{
			Employees:   getEmployees(),
			CompanyName: "Acme Corp",
			Address:     "123 Main St",
			Phone:       "555-1234",
		},
		Status: "success",
	}

	klog.Info("Preparing response: %v", response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()

	http.HandleFunc("/", handler)
	klog.Info("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
