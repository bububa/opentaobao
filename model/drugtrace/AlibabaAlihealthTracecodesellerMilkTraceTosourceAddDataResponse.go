package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
奶粉溯源-同步数据 APIResponse
alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data

奶粉溯源-同步数据
*/
type AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataResponse
}

type AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_milk_trace_tosource_add_data_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务出参true
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 操作码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 操作说明
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
