package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyhoteldetailsearchdataAPIResponse 星河-酒店详情页信息获取(新改版) API返回值
// alitrip.merchant.galaxy.hotel.detail.search.data
//
// 星河服务，获取雅高酒店详细信息，详情页新改版接口
type AlitripmerchantgalaxyhoteldetailsearchdataAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyhoteldetailsearchdataAPIResponseModel
}

// AlitripmerchantgalaxyhoteldetailsearchdataAPIResponseModel is 星河-酒店详情页信息获取(新改版) 成功返回结果
type AlitripmerchantgalaxyhoteldetailsearchdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_hotel_detail_search_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyhoteldetailsearchdataResponse `json:"result,omitempty" xml:"result,omitempty"`
}
