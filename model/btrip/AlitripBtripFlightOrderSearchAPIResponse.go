package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightOrderSearchAPIResponse 获取机票订单列表 API返回值
// alitrip.btrip.flight.order.search
//
// 第三方OA厂商获取机票订单列表
type AlitripBtripFlightOrderSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightOrderSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightOrderSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightOrderSearchAPIResponseModel).Reset()
}

// AlitripBtripFlightOrderSearchAPIResponseModel is 获取机票订单列表 成功返回结果
type AlitripBtripFlightOrderSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightOrderSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightOrderSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightOrderSearchAPIResponse)
	},
}

// GetAlitripBtripFlightOrderSearchAPIResponse 从 sync.Pool 获取 AlitripBtripFlightOrderSearchAPIResponse
func GetAlitripBtripFlightOrderSearchAPIResponse() *AlitripBtripFlightOrderSearchAPIResponse {
	return poolAlitripBtripFlightOrderSearchAPIResponse.Get().(*AlitripBtripFlightOrderSearchAPIResponse)
}

// ReleaseAlitripBtripFlightOrderSearchAPIResponse 将 AlitripBtripFlightOrderSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightOrderSearchAPIResponse(v *AlitripBtripFlightOrderSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightOrderSearchAPIResponse.Put(v)
}
