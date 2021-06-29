package xhotelofficial

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住订单取消 APIResponse
taobao.xhotel.order.official.cancel

官网信用住订单取消
*/
type TaobaoXhotelOrderOfficialCancelAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderOfficialCancelResponse
}

type TaobaoXhotelOrderOfficialCancelResponse struct {
    XMLName xml.Name `xml:"xhotel_order_official_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回提示信息
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
