package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainFlightDetailSearchAPIResponse 【商旅】机票订单详情查询 API返回值
// alitrip.btrip.supplychain.flight.detail.search
//
// 【商旅】机票订单详情查询
type AlitripBtripSupplychainFlightDetailSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainFlightDetailSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainFlightDetailSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainFlightDetailSearchAPIResponseModel).Reset()
}

// AlitripBtripSupplychainFlightDetailSearchAPIResponseModel is 【商旅】机票订单详情查询 成功返回结果
type AlitripBtripSupplychainFlightDetailSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_flight_detail_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainFlightDetailSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripSupplychainFlightDetailSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainFlightDetailSearchAPIResponse)
	},
}

// GetAlitripBtripSupplychainFlightDetailSearchAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainFlightDetailSearchAPIResponse
func GetAlitripBtripSupplychainFlightDetailSearchAPIResponse() *AlitripBtripSupplychainFlightDetailSearchAPIResponse {
	return poolAlitripBtripSupplychainFlightDetailSearchAPIResponse.Get().(*AlitripBtripSupplychainFlightDetailSearchAPIResponse)
}

// ReleaseAlitripBtripSupplychainFlightDetailSearchAPIResponse 将 AlitripBtripSupplychainFlightDetailSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainFlightDetailSearchAPIResponse(v *AlitripBtripSupplychainFlightDetailSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainFlightDetailSearchAPIResponse.Put(v)
}
