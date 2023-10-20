package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse 抽奖用户名单查询接口 API返回值
// alitrip.merchant.galaxy.wechat.data.lottery.query
//
// 抽奖用户名单查询接口
type AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponseModel
}

// AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponseModel is 抽奖用户名单查询接口 成功返回结果
type AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_data_lottery_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageableResponse `json:"result,omitempty" xml:"result,omitempty"`
}
