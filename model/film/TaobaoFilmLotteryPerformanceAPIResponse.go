package film

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmLotteryPerformanceAPIResponse 淘票票履约发放权益 API返回值
// taobao.film.lottery.performance
//
// 对外第三方合作渠道通过抽奖形式发放权益
type TaobaoFilmLotteryPerformanceAPIResponse struct {
	model.CommonResponse
	TaobaoFilmLotteryPerformanceAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilmLotteryPerformanceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilmLotteryPerformanceAPIResponseModel).Reset()
}

// TaobaoFilmLotteryPerformanceAPIResponseModel is 淘票票履约发放权益 成功返回结果
type TaobaoFilmLotteryPerformanceAPIResponseModel struct {
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

// Reset 清空结构体
func (m *TaobaoFilmLotteryPerformanceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnCode = ""
	m.ReturnErrorStackTrace = ""
	m.ReturnMessage = ""
	m.ReturnErrorOper = ""
	m.ReturnErrorSolution = ""
	m.ReturnValue = nil
}

var poolTaobaoFilmLotteryPerformanceAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilmLotteryPerformanceAPIResponse)
	},
}

// GetTaobaoFilmLotteryPerformanceAPIResponse 从 sync.Pool 获取 TaobaoFilmLotteryPerformanceAPIResponse
func GetTaobaoFilmLotteryPerformanceAPIResponse() *TaobaoFilmLotteryPerformanceAPIResponse {
	return poolTaobaoFilmLotteryPerformanceAPIResponse.Get().(*TaobaoFilmLotteryPerformanceAPIResponse)
}

// ReleaseTaobaoFilmLotteryPerformanceAPIResponse 将 TaobaoFilmLotteryPerformanceAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilmLotteryPerformanceAPIResponse(v *TaobaoFilmLotteryPerformanceAPIResponse) {
	v.Reset()
	poolTaobaoFilmLotteryPerformanceAPIResponse.Put(v)
}
