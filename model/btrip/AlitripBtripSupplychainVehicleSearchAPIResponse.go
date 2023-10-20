package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainVehicleSearchAPIResponse 【商旅】用车订单搜索 API返回值
// alitrip.btrip.supplychain.vehicle.search
//
// 【商旅】用车订单搜索
type AlitripBtripSupplychainVehicleSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainVehicleSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainVehicleSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainVehicleSearchAPIResponseModel).Reset()
}

// AlitripBtripSupplychainVehicleSearchAPIResponseModel is 【商旅】用车订单搜索 成功返回结果
type AlitripBtripSupplychainVehicleSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_vehicle_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainVehicleSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripSupplychainVehicleSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainVehicleSearchAPIResponse)
	},
}

// GetAlitripBtripSupplychainVehicleSearchAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainVehicleSearchAPIResponse
func GetAlitripBtripSupplychainVehicleSearchAPIResponse() *AlitripBtripSupplychainVehicleSearchAPIResponse {
	return poolAlitripBtripSupplychainVehicleSearchAPIResponse.Get().(*AlitripBtripSupplychainVehicleSearchAPIResponse)
}

// ReleaseAlitripBtripSupplychainVehicleSearchAPIResponse 将 AlitripBtripSupplychainVehicleSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainVehicleSearchAPIResponse(v *AlitripBtripSupplychainVehicleSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainVehicleSearchAPIResponse.Put(v)
}
