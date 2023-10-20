package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityMarketingPopupAPIResponse 营销弹屏 API返回值
// alitrip.merchant.galaxy.activity.marketing.popup
//
// 星河=活动营销弹屏
type AlitripMerchantGalaxyActivityMarketingPopupAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyActivityMarketingPopupAPIResponseModel
}

// AlitripMerchantGalaxyActivityMarketingPopupAPIResponseModel is 营销弹屏 成功返回结果
type AlitripMerchantGalaxyActivityMarketingPopupAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_marketing_popup_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyActivityMarketingPopupResponse `json:"result,omitempty" xml:"result,omitempty"`
}
