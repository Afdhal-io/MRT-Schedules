# 🚇 MRT Schedules API

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-00A86B?style=for-the-badge)
![REST API](https://img.shields.io/badge/API-REST-orange?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Portfolio_Project-blue?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-lightgrey?style=for-the-badge)

Backend REST API built with **Golang** and **Gin Framework** to provide MRT Jakarta station data using clean modular architecture.

---

## ✨ Highlights

✅ Built with Go (Golang)  
✅ Gin Gonic Framework  
✅ RESTful API Design  
✅ Modular Clean Structure  
✅ External API Integration  
✅ JSON Response Handling  
✅ Error Handling  
✅ Portfolio Ready Project

---

## 📌 Project Purpose

This project was created to demonstrate backend engineering fundamentals using Golang, including:

- Route management
- Service layer architecture
- External data fetching
- Structured JSON responses
- Maintainable project structure

---

## 🛠 Tech Stack

| Technology | Usage |
|-----------|------|
| Go | Main programming language |
| Gin | HTTP web framework |
| REST API | API architecture |
| JSON | Data exchange |
| GitHub | Version control |

---

## 📂 Folder Structure

```bash
MRT-Schedules/
├── main.go
├── go.mod
├── README.md
└── modules/
    └── station/
        ├── router.go
        ├── service.go
        ├── dto.go
        └── common/
            └── client/
```

### 1️⃣ Clone Repository

```bash
git clone https://github.com/YOUR_USERNAME/YOUR_REPOSITORY.git
cd YOUR_REPOSITORY
```

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

### 3️⃣ Run Application

```bash
go run main.go
```

### Server will run on 
```bash
http://localhost:8080
```

### 🔌 API Endpoints

### Get All Stations (on postman)
```bash
http://localhost:8080/v1/api/stations
```

### Example Response
```bash
{
  "success": true,
  "message": "Successfully get all station",
  "data": [
    {
      "id": 1,
      "name": "Bundaran HI"
    },
    {
      "id": 2,
      "name": "Dukuh Atas"
    }
  ]
}
```

### ⚠️ Note

This project originally consumed public MRT Jakarta external endpoint data.

Some upstream public endpoints may no longer be available or may return deprecated responses.

This repository is maintained for portfolio purposes to showcase backend development skills using Golang.

### 👨‍💻 Author
Afdhal-io
