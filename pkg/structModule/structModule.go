package structModule

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ValAPIKey struct {
	ID     string `json:"ID"`
	APIKey string `json:"APIKey"`
}
type EBookInfo struct {
	No      int    `json:"No"`
	Name    string `json:"Name"`
	ISBN    string `json:"ISBN"`
	Forsale string `json:"Forsale"`
	Price   int    `json:"Price"`
}

type SingleKey struct {
	Key struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"key"`
}

type MultipleKeys struct {
	/* Keys []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"keys"` */
	Keys []Keys `json:"keys"`
}

type Keys struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ValueStruct_req struct {
	Value string `json:"value"`
}

type NameStruct_req struct {
	Name string `json:"name"`
}

type Trans_struct struct {
	Translations []Translations `json:"translations"`
	/* 	Translations []struct {
		ID     int    `json:"id"`
		KeyId  string `json:"keyId"`
		Locale string `json:"locale"`
		Value  string `json:"value"`
	} `json:"translations"` */
}
type Translations struct {
	ID     int    `json:"id"`
	KeyId  string `json:"keyId"`
	Locale string `json:"locale"`
	Value  string `json:"value"`
}
type MessageStruct struct {
	Q string `json:"q"`
}

type DetectResult struct {
	Data struct {
		Detections []struct {
			Confidence float64 `json:"confidence"`
			IsReliable bool    `json:"isReliable"`
			Language   string  `json:"language"`
		} `json:"detections"`
	} `json:"data"`
}

type LocaleStruct_req struct {
	Locale string `json:"locale"`
}
