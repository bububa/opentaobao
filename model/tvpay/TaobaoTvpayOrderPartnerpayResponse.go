package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付第三方支付订单 APIResponse
taobao.tvpay.order.partnerpay

tv支付第三方发起并支付订单（使用设备授权）
*/
type TaobaoTvpayOrderPartnerpayAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayOrderPartnerpayResponse
}

type TaobaoTvpayOrderPartnerpayResponse struct {
    XMLName xml.Name `xml:"tvpay_order_partnerpay_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
