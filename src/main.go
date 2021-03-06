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

	"../pkg/validationModule"
	"github.com/julienschmidt/httprouter"
)

// 8. Language detect
func LanguagDetect(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	t_message := queryValues.Get("message")
	respondResult := controllerModule.LanguagDetectController(t_message)
	fmt.Println("respondResult", respondResult)

	//fmt.Fprintf(w, "hello, %s!\n", queryValues.Get("name"))
	/* test, err := ioutil.ReadAll(r.Body)
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
	} */
	//fmt.Println("temp ", temp)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	onlyRocale := structModule.LocaleStruct_req{respondResult.Data.Detections[0].Language}
	temp, err := json.Marshal(onlyRocale)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println()
	w.Write(temp)
}

// 7.키의 특정 언어 번역 수정하기
func UpdateTransLocaleValue(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	test, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var testAuth structModule.ValueStruct_req
	valiresult := ps.ByName("locale")
	if !validationModule.CheckLocaleValue(valiresult) {
		result := structModule.FailRespond{false, "check locale . it will be ko or en or ja"}
		temp, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusForbidden)
		w.Write(temp)
		return
	}
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
	valiresult := ps.ByName("locale")
	if !validationModule.CheckLocaleValue(valiresult) {
		result := structModule.FailRespond{false, "check locale . it will be ko or en or ja"}
		temp, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusForbidden)
		w.Write(temp)
		return
	}
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
	valiresult := ps.ByName("locale")
	if !validationModule.CheckLocaleValue(valiresult) {
		result := structModule.FailRespond{false, "check locale . it will be ko or en or ja"}
		temp, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusForbidden)
		w.Write(temp)
		return
	}
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
	/* 	test, err := ioutil.ReadAll(r.Body)
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
	   	w.Write(temp) */

	test, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var testAuth structModule.NameStruct_req
	var temp []byte
	err = json.Unmarshal(test, &testAuth)

	if controllerModule.CheckUniquName(testAuth.Name) > 0 {
		result := structModule.FailRespond{false, "check name . it will be unique value"}
		temp, err = json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusForbidden)
		w.Write(temp)
		return
	}

	if validationModule.IsSAlphabetAndDot(testAuth.Name) {
		respondResult := controllerModule.AppendKeys(testAuth.Name)
		temp, err = json.Marshal(respondResult)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		result := structModule.FailRespond{false, "check value . it will be small alphabet and dot"}
		temp, err = json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusForbidden)

	}

	// 만약에 추가하는 키에 영문자와 dot 외에 다른게 있으면 에러를 응답하고 해당 룰에 대해서 전달

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
	var temp []byte
	err = json.Unmarshal(test, &testAuth)

	if controllerModule.CheckUniquName(testAuth.Name) > 0 {
		result := structModule.FailRespond{false, "check name . it will be unique value"}
		temp, err = json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusForbidden)
		w.Write(temp)
		return
	}

	if validationModule.IsSAlphabetAndDot(testAuth.Name) {
		respondResult := controllerModule.AppendKeys(testAuth.Name)
		temp, err = json.Marshal(respondResult)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		result := structModule.FailRespond{false, "check value . it will be small alphabet and dot"}
		temp, err = json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusForbidden)

	}

	// 만약에 추가하는 키에 영문자와 dot 외에 다른게 있으면 에러를 응답하고 해당 룰에 대해서 전달

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)
}

// 1. 모든 키 가져오기
func getAllKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	t_message := queryValues.Get("name")
	respondResult := controllerModule.AllKeys(t_message)
	temp, err := json.Marshal(respondResult)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)
	//return
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
	router.GET("/language_detect", LanguagDetect)

	log.Fatal(http.ListenAndServe(":8090", router))
	//fmt.Println(http.ListenAndServe(":8090", router))

}
