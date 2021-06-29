package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消叫车 APIResponse
alibaba.happytrip.taxi.order.cancel

取消叫车订单,行程中的订单不能取消
*/
type AlibabaHappytripTaxiOrderCancelAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiOrderCancelResponse
}

type AlibabaHappytripTaxiOrderCancelResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`

    
    // 错误信息
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
    // 订单取消结果
    
    Data   *OrderCancelResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
