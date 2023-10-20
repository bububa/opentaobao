package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainHotelSearchAPIResponse 【商旅】酒店订单查询 API返回值
// alitrip.btrip.supplychain.hotel.search
//
// 【商旅】酒店订单查询
type AlitripBtripSupplychainHotelSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainHotelSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainHotelSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainHotelSearchAPIResponseModel).Reset()
}

// AlitripBtripSupplychainHotelSearchAPIResponseModel is 【商旅】酒店订单查询 成功返回结果
type AlitripBtripSupplychainHotelSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_hotel_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainHotelSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripSupplychainHotelSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainHotelSearchAPIResponse)
	},
}

// GetAlitripBtripSupplychainHotelSearchAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainHotelSearchAPIResponse
func GetAlitripBtripSupplychainHotelSearchAPIResponse() *AlitripBtripSupplychainHotelSearchAPIResponse {
	return poolAlitripBtripSupplychainHotelSearchAPIResponse.Get().(*AlitripBtripSupplychainHotelSearchAPIResponse)
}

// ReleaseAlitripBtripSupplychainHotelSearchAPIResponse 将 AlitripBtripSupplychainHotelSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainHotelSearchAPIResponse(v *AlitripBtripSupplychainHotelSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainHotelSearchAPIResponse.Put(v)
}
