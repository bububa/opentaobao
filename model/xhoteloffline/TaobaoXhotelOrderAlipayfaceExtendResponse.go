package xhoteloffline

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住订单延住接口 APIResponse
taobao.xhotel.order.alipayface.extend

信用住订单延住接口，用于将已有的信用住支付订单checkOut和订单金额等更新
*/
type TaobaoXhotelOrderAlipayfaceExtendAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderAlipayfaceExtendResponse
}

type TaobaoXhotelOrderAlipayfaceExtendResponse struct {
    XMLName xml.Name `xml:"xhotel_order_alipayface_extend_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
