package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
奶粉溯源-同步数据 APIRequest
alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data

奶粉溯源-同步数据
*/
type AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest struct {
    model.Params

    // 奶粉品牌ID
    entId   string 

    // 奶粉数据
    jsonStr   string 

}

func NewAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest() *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest{
    return &AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data"
}

func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) SetEntId(entId string) error {
    r.entId = entId
    r.Set("ent_id", entId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) GetEntId() string {
    return r.entId
}

func (r *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) SetJsonStr(jsonStr string) error {
    r.jsonStr = jsonStr
    r.Set("json_str", jsonStr)
    return nil
}

func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) GetJsonStr() string {
    return r.jsonStr
}

