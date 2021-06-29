package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详情 APIResponse
alibaba.happytrip.taxi.order.get

获取订单状态及详情
*/
type AlibabaHappytripTaxiOrderGetAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiOrderGetResponse
}

type AlibabaHappytripTaxiOrderGetResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`

    
    // 错误信息
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
    // 订单获取结果
    
    Data   *OrderGetResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
