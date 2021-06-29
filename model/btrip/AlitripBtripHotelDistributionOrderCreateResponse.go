package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店分销-创建订单 APIResponse
alitrip.btrip.hotel.distribution.order.create

商旅酒店分销-创建订单
*/
type AlitripBtripHotelDistributionOrderCreateAPIResponse struct {
    model.CommonResponse
    AlitripBtripHotelDistributionOrderCreateResponse
}

type AlitripBtripHotelDistributionOrderCreateResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 创单返回结果
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
