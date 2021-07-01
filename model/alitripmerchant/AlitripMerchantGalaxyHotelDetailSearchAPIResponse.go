package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyHotelDetailSearchAPIResponse
星河-酒店详细信息搜索 API返回值
alitrip.merchant.galaxy.hotel.detail.search

星河服务=获取雅高酒店详细信息 */
type AlitripMerchantGalaxyHotelDetailSearchAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyHotelDetailSearchAPIResponseModel
}

// AlitripMerchantGalaxyHotelDetailSearchAPIResponseModel is 星河-酒店详细信息搜索 成功返回结果
type AlitripMerchantGalaxyHotelDetailSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_hotel_detail_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripMerchantGalaxyHotelDetailSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}
