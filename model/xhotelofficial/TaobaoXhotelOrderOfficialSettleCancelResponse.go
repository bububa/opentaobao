package xhotelofficial

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住取消结账 APIResponse
taobao.xhotel.order.official.settle.cancel

用于官网信用住取消结账
*/
type TaobaoXhotelOrderOfficialSettleCancelAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderOfficialSettleCancelResponse
}

type TaobaoXhotelOrderOfficialSettleCancelResponse struct {
    XMLName xml.Name `xml:"xhotel_order_official_settle_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回信息
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
