package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentCancelAPIResponse 取消结算调整单 API返回值
// tmall.service.settleadjustment.cancel
//
// 提供给服务商在对取消已经发起的结算调整单。
// 通过说明调整单ID进行结算调整单取消。
type TmallServiceSettleadjustmentCancelAPIResponse struct {
	model.CommonResponse
	TmallServiceSettleadjustmentCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServiceSettleadjustmentCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServiceSettleadjustmentCancelAPIResponseModel).Reset()
}

// TmallServiceSettleadjustmentCancelAPIResponseModel is 取消结算调整单 成功返回结果
type TmallServiceSettleadjustmentCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settleadjustment_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServiceSettleadjustmentCancelResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServiceSettleadjustmentCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServiceSettleadjustmentCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentCancelAPIResponse)
	},
}

// GetTmallServiceSettleadjustmentCancelAPIResponse 从 sync.Pool 获取 TmallServiceSettleadjustmentCancelAPIResponse
func GetTmallServiceSettleadjustmentCancelAPIResponse() *TmallServiceSettleadjustmentCancelAPIResponse {
	return poolTmallServiceSettleadjustmentCancelAPIResponse.Get().(*TmallServiceSettleadjustmentCancelAPIResponse)
}

// ReleaseTmallServiceSettleadjustmentCancelAPIResponse 将 TmallServiceSettleadjustmentCancelAPIResponse 保存到 sync.Pool
func ReleaseTmallServiceSettleadjustmentCancelAPIResponse(v *TmallServiceSettleadjustmentCancelAPIResponse) {
	v.Reset()
	poolTmallServiceSettleadjustmentCancelAPIResponse.Put(v)
}
