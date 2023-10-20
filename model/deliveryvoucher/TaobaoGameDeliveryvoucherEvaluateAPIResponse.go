package deliveryvoucher

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherEvaluateAPIResponse 卡券评价回传 API返回值
// taobao.game.deliveryvoucher.evaluate
//
// 卡券ISV回传商品评价
type TaobaoGameDeliveryvoucherEvaluateAPIResponse struct {
	model.CommonResponse
	TaobaoGameDeliveryvoucherEvaluateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoGameDeliveryvoucherEvaluateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoGameDeliveryvoucherEvaluateAPIResponseModel).Reset()
}

// TaobaoGameDeliveryvoucherEvaluateAPIResponseModel is 卡券评价回传 成功返回结果
type TaobaoGameDeliveryvoucherEvaluateAPIResponseModel struct {
	XMLName xml.Name `xml:"game_deliveryvoucher_evaluate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 操作状态
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoGameDeliveryvoucherEvaluateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ResultCode = ""
	m.IsSuccess = false
}

var poolTaobaoGameDeliveryvoucherEvaluateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoGameDeliveryvoucherEvaluateAPIResponse)
	},
}

// GetTaobaoGameDeliveryvoucherEvaluateAPIResponse 从 sync.Pool 获取 TaobaoGameDeliveryvoucherEvaluateAPIResponse
func GetTaobaoGameDeliveryvoucherEvaluateAPIResponse() *TaobaoGameDeliveryvoucherEvaluateAPIResponse {
	return poolTaobaoGameDeliveryvoucherEvaluateAPIResponse.Get().(*TaobaoGameDeliveryvoucherEvaluateAPIResponse)
}

// ReleaseTaobaoGameDeliveryvoucherEvaluateAPIResponse 将 TaobaoGameDeliveryvoucherEvaluateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoGameDeliveryvoucherEvaluateAPIResponse(v *TaobaoGameDeliveryvoucherEvaluateAPIResponse) {
	v.Reset()
	poolTaobaoGameDeliveryvoucherEvaluateAPIResponse.Put(v)
}
