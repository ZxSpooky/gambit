package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	//"strconv"
	//"strings"
	"github.com/ZxSpooky/gambit/models"
	"github.com/ZxSpooky/gambit/tools"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/ZxSpooky/gambit/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza registro de InsertCategory")
	err := DbConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO category (Categ_name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	LastInserId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}
	fmt.Println("Insert Category > Ejecucion Exitosa")
	return LastInserId, nil

}

func UpdateCategory(c models.Category) error {
	fmt.Println("Comienza registro de UpdateCategory")
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "UPDATE category SET "
	if len(c.CategName) > 0 {
		sentencia += " Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}
	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia += ", "
		}
		sentencia += "Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}
	sentencia += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Update Catefgory > Ejecucion exitosa")
	return nil
}
