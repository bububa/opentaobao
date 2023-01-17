package usergrowth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIResponse 批量发送push API返回值
// alibaba.alsc.growth.interactive.mini.game.notice.push.batch.send
//
// 批量发送push
type AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIResponse struct {
	model.CommonResponse
	AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIResponseModel
}

// AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIResponseModel is 批量发送push 成功返回结果
type AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_mini_game_notice_push_batch_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 链路id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 错误描述
	BizErrorMsg string `json:"biz_error_msg,omitempty" xml:"biz_error_msg,omitempty"`
	// 时间戳
	CurrentTimestamp int64 `json:"current_timestamp,omitempty" xml:"current_timestamp,omitempty"`
	// 接口是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
