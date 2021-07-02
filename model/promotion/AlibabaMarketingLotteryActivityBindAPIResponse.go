package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryActivityBindAPIResponse 抽奖平台奖池绑定接口 API返回值
// alibaba.marketing.lottery.activity.bind
//
// 抽奖平台奖池关联接口
type AlibabaMarketingLotteryActivityBindAPIResponse struct {
	model.CommonResponse
	AlibabaMarketingLotteryActivityBindAPIResponseModel
}

// AlibabaMarketingLotteryActivityBindAPIResponseModel is 抽奖平台奖池绑定接口 成功返回结果
type AlibabaMarketingLotteryActivityBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_marketing_lottery_activity_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关联成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 错误码
	MsgCode int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
