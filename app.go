//app.go

type App struct {
    Router  *mux.Router
    DB      *sql.DB
}
