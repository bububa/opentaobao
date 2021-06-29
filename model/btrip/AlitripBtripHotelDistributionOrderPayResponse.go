package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店分销订单支付 APIResponse
alitrip.btrip.hotel.distribution.order.pay

商旅酒店分销订单支付
*/
type AlitripBtripHotelDistributionOrderPayAPIResponse struct {
    model.CommonResponse
    AlitripBtripHotelDistributionOrderPayResponse
}

type AlitripBtripHotelDistributionOrderPayResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_pay_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否支付成功
    
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`

    
    // 支付结果返回码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 支付结果返回信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
