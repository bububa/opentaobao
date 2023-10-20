package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderDetailAPIResponse 商旅酒店API分销查询订单详情 API返回值
// alitrip.btrip.hotel.distribution.order.detail
//
// 商旅酒店API分销查询订单详情
type AlitripBtripHotelDistributionOrderDetailAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionOrderDetailAPIResponseModel
}

// AlitripBtripHotelDistributionOrderDetailAPIResponseModel is 商旅酒店API分销查询订单详情 成功返回结果
type AlitripBtripHotelDistributionOrderDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单详情接口返回结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
