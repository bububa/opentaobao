package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenSupplychainVehicleTradeAPIResponse 商旅用车交易流水接口 API返回值
// alitrip.btrip.open.supplychain.vehicle.trade
//
// 商旅用车交易流水接口
type AlitripBtripOpenSupplychainVehicleTradeAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenSupplychainVehicleTradeAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenSupplychainVehicleTradeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenSupplychainVehicleTradeAPIResponseModel).Reset()
}

// AlitripBtripOpenSupplychainVehicleTradeAPIResponseModel is 商旅用车交易流水接口 成功返回结果
type AlitripBtripOpenSupplychainVehicleTradeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_supplychain_vehicle_trade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenSupplychainVehicleTradeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripOpenSupplychainVehicleTradeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenSupplychainVehicleTradeAPIResponse)
	},
}

// GetAlitripBtripOpenSupplychainVehicleTradeAPIResponse 从 sync.Pool 获取 AlitripBtripOpenSupplychainVehicleTradeAPIResponse
func GetAlitripBtripOpenSupplychainVehicleTradeAPIResponse() *AlitripBtripOpenSupplychainVehicleTradeAPIResponse {
	return poolAlitripBtripOpenSupplychainVehicleTradeAPIResponse.Get().(*AlitripBtripOpenSupplychainVehicleTradeAPIResponse)
}

// ReleaseAlitripBtripOpenSupplychainVehicleTradeAPIResponse 将 AlitripBtripOpenSupplychainVehicleTradeAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenSupplychainVehicleTradeAPIResponse(v *AlitripBtripOpenSupplychainVehicleTradeAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenSupplychainVehicleTradeAPIResponse.Put(v)
}
