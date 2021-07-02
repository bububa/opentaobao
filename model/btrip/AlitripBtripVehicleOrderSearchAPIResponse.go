package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripVehicleOrderSearchAPIResponse 用车订单查询接口 API返回值
// alitrip.btrip.vehicle.order.search
//
// 企业获取商旅用车订单数据
type AlitripBtripVehicleOrderSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripVehicleOrderSearchAPIResponseModel
}

// AlitripBtripVehicleOrderSearchAPIResponseModel is 用车订单查询接口 成功返回结果
type AlitripBtripVehicleOrderSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_vehicle_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}
