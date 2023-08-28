package RoleModel

import (
	"crud-with-go/config"
	"crud-with-go/entities"
)

func GetAll() []entities.Role {
	rows, err := config.DB.Query("SELECT * FROM roles")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var roles []entities.Role

	for rows.Next() {
		var role entities.Role

		err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)

		if err != nil {
			panic(err)
		}

		roles = append(roles, role)
	}

	return roles
}
