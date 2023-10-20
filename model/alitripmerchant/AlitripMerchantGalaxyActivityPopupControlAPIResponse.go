package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityPopupControlAPIResponse 营销弹屏疲劳度控制 API返回值
// alitrip.merchant.galaxy.activity.popup.control
//
// 星河=营销弹屏疲劳控制
type AlitripMerchantGalaxyActivityPopupControlAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyActivityPopupControlAPIResponseModel
}

// AlitripMerchantGalaxyActivityPopupControlAPIResponseModel is 营销弹屏疲劳度控制 成功返回结果
type AlitripMerchantGalaxyActivityPopupControlAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_popup_control_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyActivityPopupControlResponse `json:"result,omitempty" xml:"result,omitempty"`
}
