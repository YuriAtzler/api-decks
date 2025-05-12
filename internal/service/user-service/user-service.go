package userservice

import (
	"database/sql"

	"github.com/nunesyan/agendamento-nails/internal/database"
	usermodel "github.com/nunesyan/agendamento-nails/internal/model/user-model"
)

func GetUserById(id int) (usermodel.User, error) {
	var user usermodel.User

	db, err := database.Connect()
	if err != nil {
		return user, err
	}
	defer db.Close()

	if err = db.QueryRow("SELECT id, name FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}

	return user, nil
}
