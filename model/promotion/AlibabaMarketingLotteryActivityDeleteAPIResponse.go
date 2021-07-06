package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryActivityDeleteAPIResponse 抽奖平台活动删除接口 API返回值
// alibaba.marketing.lottery.activity.delete
//
// 抽奖平台活动删除接口
type AlibabaMarketingLotteryActivityDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaMarketingLotteryActivityDeleteAPIResponseModel
}

// AlibabaMarketingLotteryActivityDeleteAPIResponseModel is 抽奖平台活动删除接口 成功返回结果
type AlibabaMarketingLotteryActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_marketing_lottery_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// code
	MsgCode int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// result
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
