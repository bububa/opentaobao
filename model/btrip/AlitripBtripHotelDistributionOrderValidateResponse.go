package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店API分销下单前校验 API返回值 
alitrip.btrip.hotel.distribution.order.validate

商旅酒店API分销下单前校验
*/
type AlitripBtripHotelDistributionOrderValidateAPIResponse struct {
    model.CommonResponse
    AlitripBtripHotelDistributionOrderValidateResponse
}

// 商旅酒店API分销下单前校验 成功返回结果
type AlitripBtripHotelDistributionOrderValidateResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_validate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 下单前校验结果
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
