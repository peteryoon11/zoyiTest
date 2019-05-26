package controllerModule

import (
	"fmt"

	"../dbConnectModule"
	"../structModule"
)

func AllKeys() structModule.MultipleKeys {

	dbConnectModule.Mysql_Open()

	keyArrays := dbConnectModule.GetAllKeys()

	dbConnectModule.Mysql_Close()
	return keyArrays
}
func AppendKeys(name string) structModule.MultipleKeys {
	fmt.Println(" AppendKeys name ", name)
	dbConnectModule.Mysql_Open()

	keyArrays := dbConnectModule.AppendKey(name)

	dbConnectModule.Mysql_Close()
	/* 	var result structModule.MultipleKeys
	   	result.keys = keyArrays */
	return keyArrays
}
func UpdateKey(name string, keyId string) structModule.MultipleKeys {
	fmt.Println(" AppendKeys name ", name)
	dbConnectModule.Mysql_Open()

	keyArrays := dbConnectModule.UpdateKey(name, keyId)

	dbConnectModule.Mysql_Close()
	/* 	var result structModule.MultipleKeys
	   	result.keys = keyArrays */
	return keyArrays
}
func InsertTrans(value string, keyId string, locale string) structModule.Trans_struct {
	fmt.Println(" AppendKeys name ", value)
	dbConnectModule.Mysql_Open()

	keyArrays := dbConnectModule.InsertTrans(value, keyId, locale)

	dbConnectModule.Mysql_Close()
	/* 	var result structModule.MultipleKeys
	   	result.keys = keyArrays */
	return keyArrays
}
func AllTrans(value string, keyId string) structModule.Trans_struct {
	fmt.Println(" AppendKeys name ", value)
	dbConnectModule.Mysql_Open()

	keyArrays := dbConnectModule.AllTrans(value, keyId)

	dbConnectModule.Mysql_Close()
	/* 	var result structModule.MultipleKeys
	   	result.keys = keyArrays */
	return keyArrays
}
func GetUniqueTrans(keyId string, locale string) structModule.Trans_struct {
	//fmt.Println(" AppendKeys name ", value)
	dbConnectModule.Mysql_Open()

	keyArrays := dbConnectModule.GetUniqueTrans(keyId, locale)

	dbConnectModule.Mysql_Close()
	/* 	var result structModule.MultipleKeys
	   	result.keys = keyArrays */
	return keyArrays
}
func UpdateTransLocaleValue(value string, keyId string, locale string) structModule.Trans_struct {
	fmt.Println(" AppendKeys name ", value)
	dbConnectModule.Mysql_Open()

	keyArrays := dbConnectModule.InsertTrans(value, keyId, locale)

	dbConnectModule.Mysql_Close()
	/* 	var result structModule.MultipleKeys
	   	result.keys = keyArrays */
	return keyArrays
}
