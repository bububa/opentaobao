package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIResponse 本地生活用户增长互动小游戏积分发放 API返回值
// alibaba.alsc.growth.interactive.mini.game.integral.grant
//
// 本地生活用户增长互动小游戏积分发放
type AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIResponse struct {
	model.CommonResponse
	AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIResponseModel
}

// AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIResponseModel is 本地生活用户增长互动小游戏积分发放 成功返回结果
type AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_mini_game_integral_grant_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 错误信息
	BizErrorMsg string `json:"biz_error_msg,omitempty" xml:"biz_error_msg,omitempty"`
	// 链路id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 积分发放返回对象
	Data *MiniGameGrantIntegralDto `json:"data,omitempty" xml:"data,omitempty"`
	// 接口调用是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
