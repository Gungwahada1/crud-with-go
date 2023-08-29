package UserModel

import (
	"crud-with-go/config"
	"crud-with-go/entities"
)

func GetAll() []entities.User {
	rows, err := config.DB.Query("SELECT * FROM users")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User

		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.RoleId, &user.ReligionId, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

func Create(user entities.User) bool {
	result, err := config.DB.Exec(`
	INSERT INTO users (username, email, password, role_id, religion_id, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`,
		user.Username, user.Email, user.Password, user.RoleId, user.ReligionId, user.CreatedAt, user.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func GetById(id int) entities.User {
	row := config.DB.QueryRow("SELECT * FROM users WHERE id =?", id)

	var user entities.User

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.RoleId, &user.ReligionId, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		panic(err.Error())
	}

	return user
}

func Update(id int, user entities.User) bool {
	result, err := config.DB.Exec(`
  UPDATE users
  SET username =?, email =?, password =?, role_id =?, religion_id =?, updated_at =?
  WHERE id =?`,
    user.Username, user.Email, user.Password, user.RoleId, user.ReligionId, user.UpdatedAt, id,
  )

  if err!= nil {
    panic(err)
  }

  rowsAffected, err := result.RowsAffected()
  if err!= nil {
    panic(err)
  }

  return rowsAffected > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id =?", id)

  if err!= nil {
    return err
  }

  return nil
}