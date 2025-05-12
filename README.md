# 📝 To-Do REST API in Go

This is a lightweight To-Do List REST API written in **Go (Golang)**. It allows users to add and retrieve to-do items using HTTP requests. This project demonstrates the use of Go's standard libraries to build a simple RESTful service.

## 🚀 Features

- Add a new to-do item via POST request.
- Retrieve all to-do items via GET request.
- In-memory storage (no database required).
- Simple JSON-based communication.

## 🛠️ Technologies Used

- Go (Golang)
- net/http for web server
- encoding/json for JSON parsing and encoding

## 📦 API Endpoints

### ➕ Add a To-Do
**Endpoint:** `POST /todos`  
**Headers:** `Content-Type: application/json`  
**Body:**
```json
{
  "task": "Buy groceries"
}
