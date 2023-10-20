package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivitypopupqueryAPIResponse 星河-获取雅高小程序营销抽奖首页弹窗 API返回值
// alitrip.merchant.galaxy.activity.popup.query
//
// 获取雅高微信小程序，营销抽奖首页弹窗数据。
type AlitripmerchantgalaxyactivitypopupqueryAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyactivitypopupqueryAPIResponseModel
}

// AlitripmerchantgalaxyactivitypopupqueryAPIResponseModel is 星河-获取雅高小程序营销抽奖首页弹窗 成功返回结果
type AlitripmerchantgalaxyactivitypopupqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_popup_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyactivitypopupqueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
