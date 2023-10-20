package alitripmerchant

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponseModel is 抽奖用户名单查询接口 成功返回结果
type AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_data_lottery_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageableResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse
func GetAlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse() *AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse {
	return poolAlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse.Get().(*AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse 将 AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse(v *AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse.Put(v)
}
