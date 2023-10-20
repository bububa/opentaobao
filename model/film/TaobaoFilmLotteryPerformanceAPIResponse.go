package film

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofilmlotteryperformanceAPIResponse 淘票票履约发放权益 API返回值
// taobao.film.lottery.performance
//
// 对外第三方合作渠道通过抽奖形式发放权益
type TaobaofilmlotteryperformanceAPIResponse struct {
	model.CommonResponse
	TaobaofilmlotteryperformanceAPIResponseModel
}

// TaobaofilmlotteryperformanceAPIResponseModel is 淘票票履约发放权益 成功返回结果
type TaobaofilmlotteryperformanceAPIResponseModel struct {
	XMLName xml.Name `xml:"film_lottery_performance_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// returnErrorStackTrace
	ReturnErrorStackTrace string `json:"return_error_stack_trace,omitempty" xml:"return_error_stack_trace,omitempty"`
	// 错误描述
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// returnErrorOper
	ReturnErrorOper string `json:"return_error_oper,omitempty" xml:"return_error_oper,omitempty"`
	// returnErrorSolution
	ReturnErrorSolution string `json:"return_error_solution,omitempty" xml:"return_error_solution,omitempty"`
	// 返回值
	ReturnValue *LotteryPerformanceResult `json:"return_value,omitempty" xml:"return_value,omitempty"`
}
