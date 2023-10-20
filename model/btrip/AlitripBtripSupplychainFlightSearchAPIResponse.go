package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainFlightSearchAPIResponse 【商旅】机票订单查询 API返回值
// alitrip.btrip.supplychain.flight.search
//
// 【商旅】机票订单查询
type AlitripBtripSupplychainFlightSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainFlightSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainFlightSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainFlightSearchAPIResponseModel).Reset()
}

// AlitripBtripSupplychainFlightSearchAPIResponseModel is 【商旅】机票订单查询 成功返回结果
type AlitripBtripSupplychainFlightSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_flight_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainFlightSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripSupplychainFlightSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainFlightSearchAPIResponse)
	},
}

// GetAlitripBtripSupplychainFlightSearchAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainFlightSearchAPIResponse
func GetAlitripBtripSupplychainFlightSearchAPIResponse() *AlitripBtripSupplychainFlightSearchAPIResponse {
	return poolAlitripBtripSupplychainFlightSearchAPIResponse.Get().(*AlitripBtripSupplychainFlightSearchAPIResponse)
}

// ReleaseAlitripBtripSupplychainFlightSearchAPIResponse 将 AlitripBtripSupplychainFlightSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainFlightSearchAPIResponse(v *AlitripBtripSupplychainFlightSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainFlightSearchAPIResponse.Put(v)
}
