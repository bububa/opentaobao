package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyHotelDetailSearchAPIResponse 星河-酒店详细信息搜索 API返回值
// alitrip.merchant.galaxy.hotel.detail.search
//
// 星河服务=获取雅高酒店详细信息
type AlitripMerchantGalaxyHotelDetailSearchAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyHotelDetailSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyHotelDetailSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyHotelDetailSearchAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyHotelDetailSearchAPIResponseModel is 星河-酒店详细信息搜索 成功返回结果
type AlitripMerchantGalaxyHotelDetailSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_hotel_detail_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripMerchantGalaxyHotelDetailSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyHotelDetailSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyHotelDetailSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyHotelDetailSearchAPIResponse)
	},
}

// GetAlitripMerchantGalaxyHotelDetailSearchAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyHotelDetailSearchAPIResponse
func GetAlitripMerchantGalaxyHotelDetailSearchAPIResponse() *AlitripMerchantGalaxyHotelDetailSearchAPIResponse {
	return poolAlitripMerchantGalaxyHotelDetailSearchAPIResponse.Get().(*AlitripMerchantGalaxyHotelDetailSearchAPIResponse)
}

// ReleaseAlitripMerchantGalaxyHotelDetailSearchAPIResponse 将 AlitripMerchantGalaxyHotelDetailSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyHotelDetailSearchAPIResponse(v *AlitripMerchantGalaxyHotelDetailSearchAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyHotelDetailSearchAPIResponse.Put(v)
}
