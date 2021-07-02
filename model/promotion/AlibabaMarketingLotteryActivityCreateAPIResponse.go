package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryActivityCreateAPIResponse 抽奖平台奖池创建接口 API返回值
// alibaba.marketing.lottery.activity.create
//
// 抽奖平台奖池创建接口
type AlibabaMarketingLotteryActivityCreateAPIResponse struct {
	model.CommonResponse
	AlibabaMarketingLotteryActivityCreateAPIResponseModel
}

// AlibabaMarketingLotteryActivityCreateAPIResponseModel is 抽奖平台奖池创建接口 成功返回结果
type AlibabaMarketingLotteryActivityCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_marketing_lottery_activity_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 抽奖活动
	LotteryActivity *LotteryActivityExtendDto `json:"lottery_activity,omitempty" xml:"lottery_activity,omitempty"`
}
