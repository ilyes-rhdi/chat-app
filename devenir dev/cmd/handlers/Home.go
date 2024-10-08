package handlers
import( 

	"database/sql"
	"net/http"
 
)


func Main(res http.ResponseWriter, req *http.Request){
    session, _ := store.Get(req, "session-name")
    dsn := "root:ilyesgamer2005@@tcp(127.0.0.1:3306)/db"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        http.Error(res, "Database connection error", http.StatusInternalServerError)
        return
    }
    defer db.Close()
	users, err := GetAllUsers(db)
    if err != nil {
        http.Error(res, "Error fetching users", http.StatusInternalServerError)
        return
    }
    username, ok := session.Values["username"].(string)
    isAdmin, _ := session.Values["isAdmin"].(bool)
    if !ok || username == "" {
        http.Redirect(res, req, "/login", http.StatusFound) // Rediriger si l'utilisateur n'est pas connecté
        return
    }
    data := Pagedata{
        Currentuser: User{
            Name: username,
        Isadmin: isAdmin,
        },
        Users: users,
    }


	Rendertemplates(res,"Home", data)
		
}