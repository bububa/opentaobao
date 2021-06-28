package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付查询订单状态 APIResponse
taobao.tvpay.order.query

tv支付查询订单状态
*/
type TaobaoTvpayOrderQueryAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayOrderQueryResponse
}

type TaobaoTvpayOrderQueryResponse struct {
    XMLName xml.Name `xml:"tvpay_order_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
