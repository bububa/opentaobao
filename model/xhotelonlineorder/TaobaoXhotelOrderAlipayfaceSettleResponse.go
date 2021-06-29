package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住订单结账接口 APIResponse
taobao.xhotel.order.alipayface.settle

用于离店付订单在客人离店后，发起结账以及扣款等后续动作
*/
type TaobaoXhotelOrderAlipayfaceSettleAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderAlipayfaceSettleResponse
}

type TaobaoXhotelOrderAlipayfaceSettleResponse struct {
    XMLName xml.Name `xml:"xhotel_order_alipayface_settle_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
