package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川支付完成回调 APIResponse
taobao.baichuan.payresult.query

百川支付完成回调
*/
type TaobaoBaichuanPayresultQueryAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanPayresultQueryResponse
}

type TaobaoBaichuanPayresultQueryResponse struct {
    XMLName xml.Name `xml:"baichuan_payresult_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
