package handlers
import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"regexp"
    
	_ "github.com/go-sql-driver/mysql"
)


type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Mdp string `json:"mdp"`
    Isadmin bool `json:"Isadmin"`
}
type Pagedata struct {
	Currentuser User
	Users []User 
}
func GetAllUsers(db *sql.DB) ([]User, error) {
    query := "SELECT  name, email, isAdmin FROM users"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.Name, &user.Email, &user.Isadmin)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    // Check for any error that may have occurred during iteration
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}
func DeleteUser(db *sql.DB, username string) error {
    query := "DELETE FROM users WHERE name = ?"
    _, err := db.Exec(query, username)
    if err != nil {
        return err
    }
    return nil
}
func Rendertemplates(res http.ResponseWriter,tmpl string ,data interface{}){
	t, err:= template.ParseFiles("C:\\Users\\PC\\OneDrive\\Documents\\ilyes\\devenir dev\\templates\\"+tmpl+".page.tmpl")
	if err !=nil  {
	   http.Error(res,err.Error(),http.StatusInternalServerError)
	   return
	}
	err =t.Execute(res , data)
	if err != nil {
	   http.Error(res, "Error executing template", http.StatusInternalServerError)
	   fmt.Println("Error executing template:", err)
     }
   }
   func VerifyUser(db *sql.DB, identifier, password string) (bool, bool, string) {
	   var storedPassword string
	   var isAdmin bool
	   var query string
   
	   // Check if identifier is an email or username
	   if strings.Contains(identifier, "@") {
		   query = "SELECT password, isAdmin FROM users WHERE email = ?"
	   } else {
		   query = "SELECT password, isAdmin FROM users WHERE name = ?"
	   }
   
	   // Execute the query
	   err := db.QueryRow(query, identifier).Scan(&storedPassword, &isAdmin)
   
	   // Handle case where the user is not found or other SQL errors occur
	   if err != nil {
		   if err == sql.ErrNoRows {
			   return false, false, "User not found."
		   }
		   log.Println("SQL Error:", err)
		   return false, false, "Database error."
	   }
   
	   // Compare provided password with stored password (ensure passwords are hashed)
	   if password == storedPassword {
		   return true, isAdmin, "User verified."
	   } else {
		   return false, false, "Incorrect password."
	   }
   }
   func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:ilyesgamer2005@@tcp(localhost:3306)/db")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
}
func ValidateInput(user User) (bool, string) {
	// Vérification des champs vides
	if user.Name == "" || user.Email == "" || user.Mdp == "" {
		return false, "All fields (name, email, password) are required."
	}

	// Vérification de l'email avec une expression régulière
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return false, "Invalid email format."
	}

	// Vérification de la longueur du mot de passe (ex: minimum 6 caractères)
	if len(user.Mdp) < 6 {
		return false, "Password must be at least 6 characters long."
	}

	return true, ""
}
func SanitizeInput(user *User) {
	// Supprime les espaces avant et après les champs
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Mdp = strings.TrimSpace(user.Mdp)

	// Supprime les tags HTML potentiellement dangereux (protection contre XSS)
	re := regexp.MustCompile("<.*?>")
	user.Name = re.ReplaceAllString(user.Name, "")
	user.Email = re.ReplaceAllString(user.Email, "")
	user.Mdp = re.ReplaceAllString(user.Mdp, "")
}