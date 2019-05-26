package dbConnectModule

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"../structModule"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Mysql_Open() error {
	var err error
	db, err = sql.Open("mysql", "zoyi:zoyi11@tcp(127.0.0.1:3306)/zoyi_test")
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func Mysql_Close() {
	db.Close()
	db = nil
}

func UpdateTransUniq(t_value string, t_keyId string, t_locale string) structModule.Trans_struct {

	result, err := db.Exec("UPDATE `zoyi_test`.`tb_translation` SET `value`= ? WHERE  `keyid`= ? and `locale` = ? ", t_value, t_keyId, t_locale)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	rows, err := db.Query("SELECT  id, keyid , locale, value FROM  tb_translation where keyid = ? and locale = ? ", t_keyId, t_locale)
	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	defer rows.Close()

	var (
		//	no      int
		id     string
		keyid  string
		locale string
		value  string
	)
	var tempKeys []structModule.Translations

	for rows.Next() {
		err = rows.Scan(&id, &keyid, &locale, &value)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		fmt.Println("id ", id, " keyid ", keyid, " localce", locale, " value ", value)
		//	tempCount = append(tempCount, structModule.ValAPIKey{id, AuthKey})
		//	newSeqNo = seq_no
		tempId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}

		tempKeys = append(tempKeys, structModule.Translations{tempId, keyid, locale, value})
	}
	tempResult := structModule.Trans_struct{tempKeys}
	//Mysql_Close()
	return tempResult
}
func GetUniqueTrans(t_keyId string, t_locale string) structModule.Trans_struct {

	rows, err := db.Query("SELECT  id, keyid , locale, value FROM  tb_translation where keyid = ?  and locale = ? ", t_keyId, t_locale)
	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	defer rows.Close()

	var (
		//	no      int
		id     string
		keyid  string
		locale string
		value  string
	)
	var tempKeys []structModule.Translations

	for rows.Next() {
		err = rows.Scan(&id, &keyid, &locale, &value)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		fmt.Println("id ", id, " keyid ", keyid, " localce", locale, " value ", value)
		//	tempCount = append(tempCount, structModule.ValAPIKey{id, AuthKey})
		//	newSeqNo = seq_no
		tempId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}

		tempKeys = append(tempKeys, structModule.Translations{tempId, keyid, locale, value})
	}
	tempResult := structModule.Trans_struct{tempKeys}
	//Mysql_Close()
	return tempResult
}
func AllTrans(t_value string, t_keyId string) structModule.Trans_struct {

	rows, err := db.Query("SELECT  id, keyid , locale, value FROM  tb_translation where keyid = ?  ", t_keyId)
	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	defer rows.Close()

	var (
		//	no      int
		id     string
		keyid  string
		locale string
		value  string
	)
	var tempKeys []structModule.Translations

	for rows.Next() {
		err = rows.Scan(&id, &keyid, &locale, &value)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		fmt.Println("id ", id, " keyid ", keyid, " localce", locale, " value ", value)
		//	tempCount = append(tempCount, structModule.ValAPIKey{id, AuthKey})
		//	newSeqNo = seq_no
		tempId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}

		tempKeys = append(tempKeys, structModule.Translations{tempId, keyid, locale, value})
	}
	tempResult := structModule.Trans_struct{tempKeys}
	//Mysql_Close()
	return tempResult
}
func InsertTrans(t_value string, t_keyId string, t_locale string) structModule.Trans_struct {
	result, err := db.Exec("INSERT INTO `zoyi_test`.`tb_translation` (`keyId` , `locale` , `value`) VALUES (?, ?,?) ", t_keyId, t_locale, t_value)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	rows, err := db.Query("SELECT  id, keyid , locale, value FROM  tb_translation where keyid = ? and locale = ? ", t_keyId, t_locale)
	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	defer rows.Close()

	var (
		//	no      int
		id     string
		keyid  string
		locale string
		value  string
	)
	var tempKeys []structModule.Translations

	for rows.Next() {
		err = rows.Scan(&id, &keyid, &locale, &value)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		fmt.Println("id ", id, " keyid ", keyid, " localce", locale, " value ", value)
		//	tempCount = append(tempCount, structModule.ValAPIKey{id, AuthKey})
		//	newSeqNo = seq_no
		tempId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}

		tempKeys = append(tempKeys, structModule.Translations{tempId, keyid, locale, value})
	}
	tempResult := structModule.Trans_struct{tempKeys}
	//Mysql_Close()
	return tempResult
}
func UpdateKey(t_name string, t_keyId string) structModule.MultipleKeys {

	result, err := db.Exec("UPDATE `zoyi_test`.`tb_key` SET `name`=? WHERE  `id`= ? ", t_name, t_keyId)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	rows, err := db.Query("SELECT  id, name FROM  tb_key where id = ? ", t_keyId)
	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	defer rows.Close()

	var (
		//	no      int
		id   string
		name string
	)
	var tempKeys []structModule.Keys

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		fmt.Println("id ", id, " name ", name)
		//	tempCount = append(tempCount, structModule.ValAPIKey{id, AuthKey})
		//	newSeqNo = seq_no
		tempId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}

		tempKeys = append(tempKeys, structModule.Keys{tempId, name})
	}
	tempResult := structModule.MultipleKeys{tempKeys}
	//Mysql_Close()
	return tempResult
}
func AppendKey(t_name string) structModule.MultipleKeys {

	result, err := db.Exec("INSERT INTO `zoyi_test`.`tb_key` (`name`) VALUES (?)", t_name)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	rows, err := db.Query("SELECT  id, name FROM  tb_key where name = ? ", t_name)
	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	defer rows.Close()

	var (
		//	no      int
		id   string
		name string
	)
	var tempKeys []structModule.Keys

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		fmt.Println("id ", id, " name ", name)
		//	tempCount = append(tempCount, structModule.ValAPIKey{id, AuthKey})
		//	newSeqNo = seq_no
		tempId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}

		tempKeys = append(tempKeys, structModule.Keys{tempId, name})
	}
	tempResult := structModule.MultipleKeys{tempKeys}
	//Mysql_Close()
	return tempResult
}
func GetAllKeys(t_message string) structModule.MultipleKeys {

	rows, err := db.Query("SELECT  id, name FROM  tb_key")
	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	defer rows.Close()

	var (
		//	no      int
		id   string
		name string
	)
	var tempKeys []structModule.Keys

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		fmt.Println("id ", id, " name ", name)
		//	tempCount = append(tempCount, structModule.ValAPIKey{id, AuthKey})
		//	newSeqNo = seq_no
		tempId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}

		tempKeys = append(tempKeys, structModule.Keys{tempId, name})
	}
	var tempResult structModule.MultipleKeys
	if len(t_message) > 0 {
		var keysTemp []structModule.Keys
		for _, item := range tempKeys {
			if strings.EqualFold(item.Name, t_message) {
				keysTemp = append(keysTemp, item)
				tempResult = structModule.MultipleKeys{keysTemp}
			}
		}
		//tempResult = structModule.MultipleKeys{tempKeys}
	} else {
		tempResult = structModule.MultipleKeys{tempKeys}
	}

	//Mysql_Close()
	return tempResult
}
func CheckUniquName(t_name string) int {

	rows, err := db.Query("SELECT  id, name FROM  tb_key where name = ? ", t_name)
	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	defer rows.Close()

	var (
		//	no      int
		id   string
		name string
	)
	var tempKeys []structModule.Keys

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		fmt.Println("id ", id, " name ", name)
		//	tempCount = append(tempCount, structModule.ValAPIKey{id, AuthKey})
		//	newSeqNo = seq_no
		tempId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}

		tempKeys = append(tempKeys, structModule.Keys{tempId, name})
	}
	//tempResult := structModule.MultipleKeys{tempKeys}
	//Mysql_Close()
	return len(tempKeys)

}
