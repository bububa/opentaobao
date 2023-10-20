package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelOrderSearchAPIResponse 搜索酒店订单列表 API返回值
// alitrip.btrip.hotel.order.search
//
// 企业获取商旅酒店订单数据
type AlitripBtripHotelOrderSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelOrderSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripHotelOrderSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripHotelOrderSearchAPIResponseModel).Reset()
}

// AlitripBtripHotelOrderSearchAPIResponseModel is 搜索酒店订单列表 成功返回结果
type AlitripBtripHotelOrderSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripHotelOrderSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripHotelOrderSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripHotelOrderSearchAPIResponse)
	},
}

// GetAlitripBtripHotelOrderSearchAPIResponse 从 sync.Pool 获取 AlitripBtripHotelOrderSearchAPIResponse
func GetAlitripBtripHotelOrderSearchAPIResponse() *AlitripBtripHotelOrderSearchAPIResponse {
	return poolAlitripBtripHotelOrderSearchAPIResponse.Get().(*AlitripBtripHotelOrderSearchAPIResponse)
}

// ReleaseAlitripBtripHotelOrderSearchAPIResponse 将 AlitripBtripHotelOrderSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripHotelOrderSearchAPIResponse(v *AlitripBtripHotelOrderSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripHotelOrderSearchAPIResponse.Put(v)
}
