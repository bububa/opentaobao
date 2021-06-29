package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店API分销取消订单 API返回值 
alitrip.btrip.hotel.distribution.order.cancel

商旅酒店API分销取消订单
*/
type AlitripBtripHotelDistributionOrderCancelAPIResponse struct {
    model.CommonResponse
    AlitripBtripHotelDistributionOrderCancelResponse
}

// 商旅酒店API分销取消订单 成功返回结果
type AlitripBtripHotelDistributionOrderCancelResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_cancel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 取消订单返回结果
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
