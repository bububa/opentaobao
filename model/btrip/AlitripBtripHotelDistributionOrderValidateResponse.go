package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店API分销下单前校验 APIResponse
alitrip.btrip.hotel.distribution.order.validate

商旅酒店API分销下单前校验
*/
type AlitripBtripHotelDistributionOrderValidateAPIResponse struct {
    model.CommonResponse
    AlitripBtripHotelDistributionOrderValidateResponse
}

type AlitripBtripHotelDistributionOrderValidateResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_validate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 下单前校验结果
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
