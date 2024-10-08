package handlers
import(
	
	"fmt"
	"net/http"
	 "database/sql"
	 _ "github.com/go-sql-driver/mysql"
)
var db *sql.DB



func Submit(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
        // Render the login page (e.g., HTML page)
        Rendertemplates(res, "Submit",nil)
        return
    }
    if req.Method != http.MethodPost {
        http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

	err := req.ParseForm()
    if err != nil {
        http.Error(res, "Error parsing form data", http.StatusBadRequest)
        return
    }


    // Parse form data (if you are sending URL-encoded data)
    if err := req.ParseForm(); err != nil {
        http.Error(res, "Error parsing form data", http.StatusBadRequest)
        return
    }

    // Create a User struct from form data
    user := User{
        Name:     req.FormValue("username"),
        Email:    req.FormValue("email"),
        Mdp:      req.FormValue("password"),
        Isadmin:  req.FormValue("isAdmin")=="true",
    }

    // Validate and sanitize input
    ValidateInput(user)
    SanitizeInput(&user)

    // Prepare SQL statement
    stmt, err := db.Prepare("INSERT INTO users(name, email, password, isAdmin) VALUES(?, ?, ?, ?)")
    if err != nil {
        fmt.Println("Database prepare error:", err)  // Log the actual error
        http.Error(res, "Database error", http.StatusInternalServerError)
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(user.Name, user.Email, user.Mdp ,user.Isadmin)
    if err != nil {
        fmt.Println("Database exec error:", err)  
        http.Error(res, "Failed to insert user into database", http.StatusInternalServerError)
        return
    }

    // Send success response
        http.Redirect(res, req, "/Home", http.StatusFound)
    
    
}
