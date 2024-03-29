package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainFlightCityAPIResponse 机场数据查询 API返回值
// alitrip.btrip.supplychain.flight.city
//
// 机场数据查询
type AlitripBtripSupplychainFlightCityAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainFlightCityAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainFlightCityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainFlightCityAPIResponseModel).Reset()
}

// AlitripBtripSupplychainFlightCityAPIResponseModel is 机场数据查询 成功返回结果
type AlitripBtripSupplychainFlightCityAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_flight_city_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果对象
	Result *SuggestRs `json:"result,omitempty" xml:"result,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainFlightCityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Result = nil
	m.ResultCode = 0
}

var poolAlitripBtripSupplychainFlightCityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainFlightCityAPIResponse)
	},
}

// GetAlitripBtripSupplychainFlightCityAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainFlightCityAPIResponse
func GetAlitripBtripSupplychainFlightCityAPIResponse() *AlitripBtripSupplychainFlightCityAPIResponse {
	return poolAlitripBtripSupplychainFlightCityAPIResponse.Get().(*AlitripBtripSupplychainFlightCityAPIResponse)
}

// ReleaseAlitripBtripSupplychainFlightCityAPIResponse 将 AlitripBtripSupplychainFlightCityAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainFlightCityAPIResponse(v *AlitripBtripSupplychainFlightCityAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainFlightCityAPIResponse.Put(v)
}
