package ReligionModel

import (
	"crud-with-go/config"
	"crud-with-go/entities"
)

func GetAll() []entities.Religion {
	rows, err := config.DB.Query("SELECT id,name,created_at,updated_at FROM religions")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var religions []entities.Religion

	for rows.Next() {
		var religion entities.Religion

		err := rows.Scan(&religion.Id, &religion.Name, &religion.CreatedAt, &religion.UpdatedAt)

		if err != nil {
			panic(err)
		}

		religions = append(religions, religion)
	}

	return religions
}
