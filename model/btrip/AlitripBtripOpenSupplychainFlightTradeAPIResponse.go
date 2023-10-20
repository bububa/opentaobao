package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenSupplychainFlightTradeAPIResponse 【商旅】机票交易流水查询接口 API返回值
// alitrip.btrip.open.supplychain.flight.trade
//
// 【商旅】杭州市政府机票交易流水接口查询
type AlitripBtripOpenSupplychainFlightTradeAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenSupplychainFlightTradeAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenSupplychainFlightTradeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenSupplychainFlightTradeAPIResponseModel).Reset()
}

// AlitripBtripOpenSupplychainFlightTradeAPIResponseModel is 【商旅】机票交易流水查询接口 成功返回结果
type AlitripBtripOpenSupplychainFlightTradeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_supplychain_flight_trade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenSupplychainFlightTradeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripOpenSupplychainFlightTradeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenSupplychainFlightTradeAPIResponse)
	},
}

// GetAlitripBtripOpenSupplychainFlightTradeAPIResponse 从 sync.Pool 获取 AlitripBtripOpenSupplychainFlightTradeAPIResponse
func GetAlitripBtripOpenSupplychainFlightTradeAPIResponse() *AlitripBtripOpenSupplychainFlightTradeAPIResponse {
	return poolAlitripBtripOpenSupplychainFlightTradeAPIResponse.Get().(*AlitripBtripOpenSupplychainFlightTradeAPIResponse)
}

// ReleaseAlitripBtripOpenSupplychainFlightTradeAPIResponse 将 AlitripBtripOpenSupplychainFlightTradeAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenSupplychainFlightTradeAPIResponse(v *AlitripBtripOpenSupplychainFlightTradeAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenSupplychainFlightTradeAPIResponse.Put(v)
}
