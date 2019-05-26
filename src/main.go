package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//	"../pkg/dbConnectModule"
	"../pkg/controllerModule"
	"../pkg/structModule"

	//"../pkg/validationModule"
	"github.com/julienschmidt/httprouter"
)

// 7.키의 특정 언어 번역 수정하기
func UpdateTransLocaleValue(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	test, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var testAuth structModule.ValueStruct_req

	err = json.Unmarshal(test, &testAuth)
	fmt.Println("testAuth.Value ", testAuth.Value, "key id ", ps.ByName("keyId"), " locale ", ps.ByName("locale"))
	respondResult := controllerModule.UpdateTransLocaleValue(testAuth.Value, ps.ByName("keyId"), ps.ByName("locale"))
	temp, err := json.Marshal(respondResult)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("temp ", temp)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)
}

// 6.키의 특정 언어 번역 확인하기
func getUniqueTrans(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	test, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var testAuth structModule.ValueStruct_req

	err = json.Unmarshal(test, &testAuth)
	fmt.Println("testAuth.Value ", testAuth.Value, "key id ", ps.ByName("keyId"), " locale ", ps.ByName("locale"))
	respondResult := controllerModule.GetUniqueTrans(ps.ByName("keyId"), ps.ByName("locale"))
	temp, err := json.Marshal(respondResult)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("temp ", temp)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)
}

// 5.키의 모든 번역 확인하기
func allTransGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	test, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var testAuth structModule.ValueStruct_req

	err = json.Unmarshal(test, &testAuth)
	fmt.Println("testAuth.Value ", testAuth.Value, "key id ", ps.ByName("keyId"))
	respondResult := controllerModule.AllTrans(testAuth.Value, ps.ByName("keyId"))
	temp, err := json.Marshal(respondResult)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("temp ", temp)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)
}

// 4.번역 추가하기
func insertTrans(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	test, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var testAuth structModule.ValueStruct_req

	err = json.Unmarshal(test, &testAuth)
	fmt.Println("testAuth.Value ", testAuth.Value, "key id ", ps.ByName("keyId"), " locale ", ps.ByName("locale"))
	respondResult := controllerModule.InsertTrans(testAuth.Value, ps.ByName("keyId"), ps.ByName("locale"))
	temp, err := json.Marshal(respondResult)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("temp ", temp)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)
}

// 3. 키 수정하기
func updateKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//fmt.Fprintf(w, "keyId, %s!\n", ps.ByName("keyId"))
	test, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var testAuth structModule.NameStruct_req

	err = json.Unmarshal(test, &testAuth)

	respondResult := controllerModule.UpdateKey(testAuth.Name, ps.ByName("keyId"))
	temp, err := json.Marshal(respondResult)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)
}

// 2. 키 추가하기
func appendKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	test, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var testAuth structModule.NameStruct_req

	err = json.Unmarshal(test, &testAuth)

	respondResult := controllerModule.AppendKeys(testAuth.Name)
	temp, err := json.Marshal(respondResult)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)
}
func getAllKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	respondResult := controllerModule.AllKeys()
	temp, err := json.Marshal(respondResult)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)
	//return
}

func getMyBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	test, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	//	var testmem Member
	var testAuth structModule.NameStruct_req

	err = json.Unmarshal(test, &testAuth)
	fmt.Println(testAuth)
	fmt.Println(testAuth.Name)
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	fmt.Fprintf(w, "keyId, %s!\n", ps.ByName("keyId"))
	fmt.Fprintf(w, "locale, %s!\n", ps.ByName("locale"))
}
func main() {
	router := httprouter.New()

	// 1. 모든 키 가져오기
	router.GET("/keys", getAllKey)

	// 2. 키 추가하기
	router.POST("/keys", appendKey)

	// 3. 키 수정하기
	router.PUT("/keys/:keyId", updateKey)

	//4. 번역 추가하기
	router.POST("/keys/:keyId/translations/:locale", insertTrans)

	// 5. 키의 모든 번역 확인하기
	router.GET("/keys/:keyId/translations", allTransGet)

	// 6. 키의 특정 언어 번역 확인하기
	router.GET("/keys/:keyId/translations/:locale", getUniqueTrans)

	// 7. 키의 특정 언어 번역 수정하기
	router.PUT("/keys/:keyId/translations/:locale", UpdateTransLocaleValue)

	// 8. Language detect
	router.GET("/language_detect", getMyBook)

	log.Fatal(http.ListenAndServe(":8090", router))
}
