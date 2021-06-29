package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
奶粉溯源-同步数据 API请求
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

// 初始化AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest对象
func NewAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest() *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest{
    return &AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntId Setter
// 奶粉品牌ID
func (r *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) SetEntId(entId string) error {
    r.entId = entId
    r.Set("ent_id", entId)
    return nil
}

// EntId Getter
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) GetEntId() string {
    return r.entId
}
// JsonStr Setter
// 奶粉数据
func (r *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) SetJsonStr(jsonStr string) error {
    r.jsonStr = jsonStr
    r.Set("json_str", jsonStr)
    return nil
}

// JsonStr Getter
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest) GetJsonStr() string {
    return r.jsonStr
}
