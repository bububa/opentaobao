package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyHotelListSearchAPIResponse 星河-酒店列表页搜索 API返回值
// alitrip.merchant.galaxy.hotel.list.search
//
// 星河产品=酒店列表页搜索
type AlitripMerchantGalaxyHotelListSearchAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyHotelListSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyHotelListSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyHotelListSearchAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyHotelListSearchAPIResponseModel is 星河-酒店列表页搜索 成功返回结果
type AlitripMerchantGalaxyHotelListSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_hotel_list_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageableResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyHotelListSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyHotelListSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyHotelListSearchAPIResponse)
	},
}

// GetAlitripMerchantGalaxyHotelListSearchAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyHotelListSearchAPIResponse
func GetAlitripMerchantGalaxyHotelListSearchAPIResponse() *AlitripMerchantGalaxyHotelListSearchAPIResponse {
	return poolAlitripMerchantGalaxyHotelListSearchAPIResponse.Get().(*AlitripMerchantGalaxyHotelListSearchAPIResponse)
}

// ReleaseAlitripMerchantGalaxyHotelListSearchAPIResponse 将 AlitripMerchantGalaxyHotelListSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyHotelListSearchAPIResponse(v *AlitripMerchantGalaxyHotelListSearchAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyHotelListSearchAPIResponse.Put(v)
}
