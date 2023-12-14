package bd

import (
	"database/sql"
	"fmt"

	//"strconv"
	//"strings"
	"github.com/ZxSpooky/gambit/models"
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
