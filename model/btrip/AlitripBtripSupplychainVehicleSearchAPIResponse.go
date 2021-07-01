package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainVehicleSearchAPIResponse
【商旅】用车订单搜索 API返回值
alitrip.btrip.supplychain.vehicle.search

【商旅】用车订单搜索 */
type AlitripBtripSupplychainVehicleSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainVehicleSearchAPIResponseModel
}

// AlitripBtripSupplychainVehicleSearchAPIResponseModel is 【商旅】用车订单搜索 成功返回结果
type AlitripBtripSupplychainVehicleSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_vehicle_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
