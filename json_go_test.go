package main

import (
	"fmt"
	"github.com/ChengjinWu/gojson"
	"testing"
)

func Test_json_array111(t *testing.T) {
	data := `{
    "header":{
        "appid":20,
        "log_id":"dev",
        "uid":"",
        "uname":"",
        "provider": "algodemo",
        "signid":"",
        "version":"",
        "ip":""
    },
    "response":{
        "err_no":0,
        "err_msg":"",
        "results":{
             "jd_ids":[
             {
             "jd_id":1,
             "jd_name": "go语言开发"
             },
             {
               "jd_id":2,
               "jd_name": "设计师"
             }
             ],
             "cv_sample":{
               "cv_id": 1,
               "cv_json":"abc"
             },
             "jd_sample":{
               "jd_id":1,
               "jd_json":"abc"
             }
            }
        }
}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(object.GetCoding("travel"))
	}
}
