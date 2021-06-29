package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付预下单 APIResponse
taobao.tvpay.order.precreate

tv支付预下单
*/
type TaobaoTvpayOrderPrecreateAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayOrderPrecreateResponse
}

type TaobaoTvpayOrderPrecreateResponse struct {
    XMLName xml.Name `xml:"tvpay_order_precreate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
