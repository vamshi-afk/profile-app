# 🐘 Go Profile App (PostgreSQL Edition)

A lightweight Go web application for user authentication and profile management using PostgreSQL for persistent storage. Built with Gorilla Mux for routing, Gorilla Sessions for session handling, and ready for **Render deployment**.

---

## ✨ Features

* 📝 User Registration & Login
* 🔐 Session-based Authentication (Cookies)
* 🧾 Profile Management

  * Name, Email, Bio, Hobbies, Friends
* 🛡️ Protected Routes via Middleware
* 💾 PostgreSQL Database Support
* ☁️ Render-Ready Deployment
* 📦 Modular and Clean Go Project Structure

---

## 📁 Folder Structure

```bash
.
├── database
│   └── db.go             # PostgreSQL DB logic and queries
├── handlers
│   ├── auth.go           # Handles login, register, logout
│   └── profile.go        # Handles profile view and update
├── middleware
│   └── AuthMiddleware.go # Middleware for session checking
├── templates
│   ├── auth.html         # Auth page (login/register)
│   └── profile.html      # Profile dashboard
├── utils
│   └── session.go        # Gorilla session setup
├── main.go               # Entry point
├── go.mod                # Go module file
├── go.sum                # Dependency checksums
└── README.md             # You're here
```

---

## 🚀 Getting Started (Local Dev)

### ✅ Prerequisites

* Go 1.18+
* PostgreSQL installed and running
* A PostgreSQL DB URI (e.g., `postgres://user:pass@localhost:5432/dbname?sslmode=disable`)

---

### 🔧 Setup & Run

1. **Clone the repo**

```bash
git clone https://github.com/your-username/profile-app.git
cd profile-app
```

2. **Set your environment variable**

```bash
export DATABASE_URL="postgres://user:password@localhost:5432/profiledb?sslmode=disable"
```

3. **Install dependencies**

```bash
go mod tidy
```

4. **Run the app**

```bash
go run main.go
```

5. Open browser:
   👉 [http://localhost:8080](http://localhost:8080)

---

## ☁️ Deployment (Render)

1. Go to [Render.com](https://render.com/)
2. Create a **Web Service**:

   * Environment: **Go**
   * Build Command: `go build -o main .`
   * Start Command: `./main`
3. Add a **PostgreSQL Database** via Render
4. Copy the DB URL and add it as `DATABASE_URL` in your Render app's environment variables
5. Deploy 🚀

---

## 🔐 Session Middleware

The `/profile` and `/logout` routes are protected using session-based middleware:

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

## 🌐 Environment Variables

| Key            | Required | Description                        |
| -------------- | -------- | ---------------------------------- |
| `DATABASE_URL` | ✅        | PostgreSQL connection string (URI) |

Example:

```
DATABASE_URL=postgres://user:pass@localhost:5432/profiledb?sslmode=disable
```

---

## ✅ Features in Master (PostgreSQL Branch)

* Full support for persistent PostgreSQL storage
* Compatible with Render cloud deployment
* Auto table creation (if not exists)
* Clean error logging and session handling
* Safe password hashing with `bcrypt`

---

## 📌 Future Enhancements

* Profile picture uploads
* Flash messages (alerts)
* Rate limiting for brute-force protection
* Admin dashboard
* Docker support

---

## 🙌 Built With

* [Go](https://go.dev/)
* [PostgreSQL](https://www.postgresql.org/)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [Gorilla Sessions](https://github.com/gorilla/sessions)
* [Bootstrap 5](https://getbootstrap.com/)
* [Render](https://render.com/)

---
