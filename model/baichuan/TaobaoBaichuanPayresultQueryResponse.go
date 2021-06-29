package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川支付完成回调 API返回值 
taobao.baichuan.payresult.query

百川支付完成回调
*/
type TaobaoBaichuanPayresultQueryAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanPayresultQueryResponse
}

// 百川支付完成回调 成功返回结果
type TaobaoBaichuanPayresultQueryResponse struct {
    XMLName xml.Name `xml:"baichuan_payresult_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
