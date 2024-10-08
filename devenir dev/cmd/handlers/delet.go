package handlers
import ( 

	"database/sql"
	"net/http"


	

)
func DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodPost {
        http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Récupérer le nom de l'utilisateur à supprimer à partir des paramètres de la requête
    username := req.FormValue("username")
    if username == "" {
        http.Error(res, "Username not provided", http.StatusBadRequest)
        return
    }

    // Connexion à la base de données
    dsn := "root:ilyesgamer2005@@tcp(127.0.0.1:3306)/db"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        http.Error(res, "Database connection error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Supprimer l'utilisateur
    err = DeleteUser(db, username)
    if err != nil {
        http.Error(res, "Error deleting user", http.StatusInternalServerError)
        return
    }

    http.Redirect(res, req, "/Home", http.StatusFound)
}
