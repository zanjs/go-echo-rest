package utils

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/zanjs/go-echo-rest/app/models"
	"github.com/zanjs/go-echo-rest/config"
)

// Parameter 请求参数
func Parameter(method string, qmProduct models.QMProduct) models.QMRequest {
	var qmparam models.QMParameter
	var Config = config.Config

	var qm = Config.QM
	var secret = qm.Secret
	var apiURL = qm.APIURL
	qmparam.APPKey = qm.AppKey
	qmparam.CustomerID = qm.CustomerID
	qmparam.Format = qm.Format
	qmparam.Method = method
	qmparam.SignMethod = qm.SignMethod
	qmparam.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	qmparam.Version = qm.Version

	fmt.Println(qmparam)

	// APPKey     string `json:"app_key"`
	// CustomerID string `json:"customerid"`
	// Format     string `json:"format"`
	// Method     string `json:"method"`
	// SignMethod string `json:"sign_method"`
	// Timestamp  string `json:"timestamp"`
	// Version    string `json:"v"`
	var str = ""
	var strURL = ""
	// var str = "app_key=" + qm.AppKey + "&customerid=" + qm.CustomerID + "&Format=" + qm.Format
	// str = str + "&method=" + method + "&sign_method=" + qm.SignMethod + "&timestamp=" + qmparam.Timestamp
	// str = str + "&v=" + qm.Version
	value := reflect.ValueOf(qmparam)
	typ := reflect.TypeOf(qmparam)
	for i := 0; i < typ.NumField(); i++ {
		k := typ.Field(i).Tag.Get("json")
		vi := value.Field(i).Interface()
		var v string
		v = vi.(string)
		str += k + v
		strURL += "&" + k + "=" + v
		fmt.Println(typ.Field(i).Name, typ.Field(i).Tag.Get("json"), value.Field(i).Interface())
	}

	fmt.Println(str)
	fmt.Println("strURL:", strURL)

	value2 := reflect.ValueOf(qmProduct)
	typ2 := reflect.TypeOf(qmProduct)

	var xml = "<?xml version='1.0' encoding='UTF-8'?><request><criteriaList><criteria>"

	for y := 0; y < typ2.NumField(); y++ {
		k := typ2.Field(y).Tag.Get("json")
		vi := value2.Field(y).Interface()
		var v string
		v = vi.(string)

		xml += "<" + k + ">" + v + "</" + k + ">"

		fmt.Println(typ2.Field(y).Name, typ2.Field(y).Tag.Get("json"), value2.Field(y).Interface())
	}

	xml += "</criteria></criteriaList></request>"

	fmt.Println("xml:", xml)

	restr := secret + str + xml + secret

	fmt.Println("restr: ", restr)

	// md5Ctx := md5.New()
	data := []byte(restr)
	restrMD5 := md5.Sum(data)
	restrMD52 := fmt.Sprintf("%x", restrMD5)

	sign := strings.ToUpper(restrMD52)
	fmt.Println("restr: md5", restrMD52)
	fmt.Println("restr: md5 大写", sign)

	strURL = apiURL + "?sign=" + sign + strURL

	fmt.Println("strURL: sign", strURL)

	var qmRequest models.QMRequest

	u, _ := url.Parse(strURL)
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode

	qmRequest.URL = u.String()
	qmRequest.Body = xml

	return qmRequest
}

func ReqXML() {

}
