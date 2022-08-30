package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityDrawParticipateAPIResponse 星河-幸运抽奖活动参与 API返回值
// alitrip.merchant.galaxy.activity.draw.participate
//
// 雅高小程序幸运抽奖活动页面用户进行抽奖，根据活动规则返回抽奖结果
type AlitripMerchantGalaxyActivityDrawParticipateAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyActivityDrawParticipateAPIResponseModel
}

// AlitripMerchantGalaxyActivityDrawParticipateAPIResponseModel is 星河-幸运抽奖活动参与 成功返回结果
type AlitripMerchantGalaxyActivityDrawParticipateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_draw_participate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyActivityDrawParticipateResponse `json:"result,omitempty" xml:"result,omitempty"`
}
