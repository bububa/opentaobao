package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyHotelDetailSearchDataAPIResponse 星河-酒店详情页信息获取(新改版) API返回值
// alitrip.merchant.galaxy.hotel.detail.search.data
//
// 星河服务，获取雅高酒店详细信息，详情页新改版接口
type AlitripMerchantGalaxyHotelDetailSearchDataAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyHotelDetailSearchDataAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyHotelDetailSearchDataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyHotelDetailSearchDataAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyHotelDetailSearchDataAPIResponseModel is 星河-酒店详情页信息获取(新改版) 成功返回结果
type AlitripMerchantGalaxyHotelDetailSearchDataAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_hotel_detail_search_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyHotelDetailSearchDataResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyHotelDetailSearchDataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyHotelDetailSearchDataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyHotelDetailSearchDataAPIResponse)
	},
}

// GetAlitripMerchantGalaxyHotelDetailSearchDataAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyHotelDetailSearchDataAPIResponse
func GetAlitripMerchantGalaxyHotelDetailSearchDataAPIResponse() *AlitripMerchantGalaxyHotelDetailSearchDataAPIResponse {
	return poolAlitripMerchantGalaxyHotelDetailSearchDataAPIResponse.Get().(*AlitripMerchantGalaxyHotelDetailSearchDataAPIResponse)
}

// ReleaseAlitripMerchantGalaxyHotelDetailSearchDataAPIResponse 将 AlitripMerchantGalaxyHotelDetailSearchDataAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyHotelDetailSearchDataAPIResponse(v *AlitripMerchantGalaxyHotelDetailSearchDataAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyHotelDetailSearchDataAPIResponse.Put(v)
}
