//app.go

type App struct {
    Router  *mux.Router
    DB      *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {}

fund (a *App) Run(addr string) {}
