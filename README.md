# 👤 Go Profile App

A simple web application built with Go that allows users to register, log in, and manage their profile. It uses Gorilla Mux for routing, Gorilla Sessions for session management, and SQLite for persistent storage. The UI is styled with Bootstrap 5 for a clean and responsive experience.

---

## ✨ Features

- 🔐 User registration, login, and logout
- 🔄 Cookie-based session management
- 👤 Profile page with editable fields:
  - Name
  - Email
  - Bio
  - Hobbies
  - Friends
- 🛡️ Middleware-protected routes
- 💾 SQLite3 for persistent storage
- 🎨 Modern Bootstrap-based UI
- 📦 Modular Go project structure

---

## 📁 Folder Structure
```bash
.
├── database
│   ├── db.go            # Database connection & queries
│   └── users.db         # SQLite3 database file
├── handlers
│   ├── auth.go          # Handlers for register/login
│   └── profile.go       # Handler for profile view/update
├── middleware
│   └── AuthMiddleware.go # Middleware to protect routes
├── templates
│   ├── auth.html        # Login/Register HTML
│   └── profile.html     # Profile page HTML
├── utils
│   └── session.go       # Session store setup
├── main.go              # App entry point
├── go.mod               # Go module file
├── go.sum               # Dependency checksums
├── .gitignore           # Git ignore rules
└── README.md            # This file
```
---

## 🚀 Getting Started

### 📦 Prerequisites

- Go 1.18 or higher
- Git installed
- SQLite (no setup needed, auto-creates DB file)

---

### 🛠️ Installation & Run

1. **Clone the repo**
```bash
git clone https://github.com/your-username/go-profile-app.git
cd go-profile-app
````

2. **Install dependencies**

```bash
go mod tidy
```

3. **Run the server**

```bash
go run main.go
```

4. **Open your browser**
   Visit: [http://localhost:8080](http://localhost:8080)

---

## 🧪 How to Use

* Register with a unique username and password
* Log in to access your profile
* Edit your profile fields and save
* Logout to clear session

---

## 🔐 Middleware Explained

The app uses a simple auth middleware to protect private routes like `/profile` and `/logout`. It checks if the session contains a valid username. If not, the user is redirected to the home page.

```go
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := utils.Store.Get(r, "session-name")
		username, ok := session.Values["username"].(string)
		if !ok || username == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
```

---

## 🙌 Built With

* [Go](https://go.dev/)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [Gorilla Sessions](https://github.com/gorilla/sessions)
* [Bootstrap](https://getbootstrap.com/)
* [SQLite](https://www.sqlite.org/index.html)

