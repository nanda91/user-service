package postgresql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// Info ...
type Info struct {
	Debug    bool
	Hostname string
	Database string
	Username string
	Password string
	Port     string
}

// Connect ...
func (i *Info) Connect() (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		i.Hostname,
		i.Username,
		i.Database,
		i.Password)

	db, err := gorm.Open("postgres", connString)

	if err != nil {
		panic(err)
	}

	return db, err
}
