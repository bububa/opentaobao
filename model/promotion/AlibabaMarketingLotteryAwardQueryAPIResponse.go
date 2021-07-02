package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryAwardQueryAPIResponse 抽奖平台查询可用奖品接口 API返回值
// alibaba.marketing.lottery.award.query
//
// 抽奖平台查询可用奖品接口
type AlibabaMarketingLotteryAwardQueryAPIResponse struct {
	model.CommonResponse
	AlibabaMarketingLotteryAwardQueryAPIResponseModel
}

// AlibabaMarketingLotteryAwardQueryAPIResponseModel is 抽奖平台查询可用奖品接口 成功返回结果
type AlibabaMarketingLotteryAwardQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_marketing_lottery_award_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *LotteryAwardInstResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// code
	MsgCode int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// msg
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
