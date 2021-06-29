package xhotelofficial

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住取消结账 API返回值 
taobao.xhotel.order.official.settle.cancel

用于官网信用住取消结账
*/
type TaobaoXhotelOrderOfficialSettleCancelAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderOfficialSettleCancelResponse
}

// 官网信用住取消结账 成功返回结果
type TaobaoXhotelOrderOfficialSettleCancelResponse struct {
    XMLName xml.Name `xml:"xhotel_order_official_settle_cancel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回信息
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
