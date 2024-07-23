package djangolang

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type ServerConfig struct {
	DB  *sql.DB
	RTR *mux.Router
}

type NullString struct {
	sql.NullString
}
