package xhoteloffline

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下信用住订单取消 APIResponse
taobao.xhotel.order.alipayface.cancel

提供给卖家进行线下信用住的订单取消。此接口仅仅支持线下信用住订单的取消
*/
type TaobaoXhotelOrderAlipayfaceCancelAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderAlipayfaceCancelResponse
}

type TaobaoXhotelOrderAlipayfaceCancelResponse struct {
    XMLName xml.Name `xml:"xhotel_order_alipayface_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回描述
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
