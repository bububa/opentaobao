package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityFatigueAPIResponse 营销抽奖-弹窗疲劳度控制 API返回值
// alitrip.merchant.galaxy.activity.fatigue
//
// 星河产品-营销抽奖首页弹窗疲劳度控制服务
type AlitripMerchantGalaxyActivityFatigueAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyActivityFatigueAPIResponseModel
}

// AlitripMerchantGalaxyActivityFatigueAPIResponseModel is 营销抽奖-弹窗疲劳度控制 成功返回结果
type AlitripMerchantGalaxyActivityFatigueAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_fatigue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyActivityFatigueResponse `json:"result,omitempty" xml:"result,omitempty"`
}
