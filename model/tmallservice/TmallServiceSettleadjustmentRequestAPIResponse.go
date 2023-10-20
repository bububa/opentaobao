package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentRequestAPIResponse 创建结算调整单 API返回值
// tmall.service.settleadjustment.request
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明工单ID，调整费用值，调整原因进行新建结算调整单。
type TmallServiceSettleadjustmentRequestAPIResponse struct {
	model.CommonResponse
	TmallServiceSettleadjustmentRequestAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServiceSettleadjustmentRequestAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServiceSettleadjustmentRequestAPIResponseModel).Reset()
}

// TmallServiceSettleadjustmentRequestAPIResponseModel is 创建结算调整单 成功返回结果
type TmallServiceSettleadjustmentRequestAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settleadjustment_request_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServiceSettleadjustmentRequestResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServiceSettleadjustmentRequestAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServiceSettleadjustmentRequestAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentRequestAPIResponse)
	},
}

// GetTmallServiceSettleadjustmentRequestAPIResponse 从 sync.Pool 获取 TmallServiceSettleadjustmentRequestAPIResponse
func GetTmallServiceSettleadjustmentRequestAPIResponse() *TmallServiceSettleadjustmentRequestAPIResponse {
	return poolTmallServiceSettleadjustmentRequestAPIResponse.Get().(*TmallServiceSettleadjustmentRequestAPIResponse)
}

// ReleaseTmallServiceSettleadjustmentRequestAPIResponse 将 TmallServiceSettleadjustmentRequestAPIResponse 保存到 sync.Pool
func ReleaseTmallServiceSettleadjustmentRequestAPIResponse(v *TmallServiceSettleadjustmentRequestAPIResponse) {
	v.Reset()
	poolTmallServiceSettleadjustmentRequestAPIResponse.Put(v)
}
