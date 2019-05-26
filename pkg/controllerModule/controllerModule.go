package controllerModule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	var keyArrays structModule.MultipleKeys
	/* if validationModule.IsSAlphabetAndDot(name) {

	} else {

	} */
	keyArrays = dbConnectModule.AppendKey(name)

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

	keyArrays := dbConnectModule.UpdateTransUniq(value, keyId, locale)

	dbConnectModule.Mysql_Close()
	/* 	var result structModule.MultipleKeys
	   	result.keys = keyArrays */
	return keyArrays
}
func LanguagDetectController(message string) structModule.DetectResult {

	url := "https://ws.detectlanguage.com"

	path := "/0.2/detect"

	temp := structModule.MessageStruct{message}
	var result structModule.DetectResult
	fmt.Println("temp ", temp)
	temp_result := HttpClient(url+path, temp)
	json.Unmarshal(temp_result, &result)
	return result
}

// https://stackoverflow.com/questions/51452148/how-can-i-make-a-request-with-a-bearer-token-in-go
func HttpClient(urlPath string, respondUser structModule.MessageStruct) []byte {
	pbytes, _ := json.Marshal(respondUser)
	buff := bytes.NewBuffer(pbytes)
	var bearer = "Bearer 56c270e9a8e381b428c0e64314651c4d"

	req, err := http.NewRequest("POST", urlPath, buff)
	req.Header.Add("User-Agent", "TestAgent")
	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("Authorization", bearer)

	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)
	return bytes
}
func CheckUniquName(t_name string) int {

	dbConnectModule.Mysql_Open()

	keyArrays := dbConnectModule.CheckUniquName(t_name)
	fmt.Println("keyArrays ", keyArrays)
	dbConnectModule.Mysql_Close()
	return keyArrays
}
