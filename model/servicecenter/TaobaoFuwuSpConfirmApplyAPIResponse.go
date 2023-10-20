package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFuwuSpConfirmApplyAPIResponse 内购服务确认单申请接口 API返回值
// taobao.fuwu.sp.confirm.apply
//
// isv能通过该接口发起确认申请单
type TaobaoFuwuSpConfirmApplyAPIResponse struct {
	model.CommonResponse
	TaobaoFuwuSpConfirmApplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFuwuSpConfirmApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFuwuSpConfirmApplyAPIResponseModel).Reset()
}

// TaobaoFuwuSpConfirmApplyAPIResponseModel is 内购服务确认单申请接口 成功返回结果
type TaobaoFuwuSpConfirmApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"fuwu_sp_confirm_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的是服务市场的确认单ID
	ApplyResult int64 `json:"apply_result,omitempty" xml:"apply_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFuwuSpConfirmApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApplyResult = 0
}

var poolTaobaoFuwuSpConfirmApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFuwuSpConfirmApplyAPIResponse)
	},
}

// GetTaobaoFuwuSpConfirmApplyAPIResponse 从 sync.Pool 获取 TaobaoFuwuSpConfirmApplyAPIResponse
func GetTaobaoFuwuSpConfirmApplyAPIResponse() *TaobaoFuwuSpConfirmApplyAPIResponse {
	return poolTaobaoFuwuSpConfirmApplyAPIResponse.Get().(*TaobaoFuwuSpConfirmApplyAPIResponse)
}

// ReleaseTaobaoFuwuSpConfirmApplyAPIResponse 将 TaobaoFuwuSpConfirmApplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFuwuSpConfirmApplyAPIResponse(v *TaobaoFuwuSpConfirmApplyAPIResponse) {
	v.Reset()
	poolTaobaoFuwuSpConfirmApplyAPIResponse.Put(v)
}
