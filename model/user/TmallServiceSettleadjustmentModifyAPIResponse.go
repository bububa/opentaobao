package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentModifyAPIResponse 修改结算调整单 API返回值
// tmall.service.settleadjustment.modify
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
type TmallServiceSettleadjustmentModifyAPIResponse struct {
	model.CommonResponse
	TmallServiceSettleadjustmentModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServiceSettleadjustmentModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServiceSettleadjustmentModifyAPIResponseModel).Reset()
}

// TmallServiceSettleadjustmentModifyAPIResponseModel is 修改结算调整单 成功返回结果
type TmallServiceSettleadjustmentModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settleadjustment_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServiceSettleadjustmentModifyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServiceSettleadjustmentModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServiceSettleadjustmentModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentModifyAPIResponse)
	},
}

// GetTmallServiceSettleadjustmentModifyAPIResponse 从 sync.Pool 获取 TmallServiceSettleadjustmentModifyAPIResponse
func GetTmallServiceSettleadjustmentModifyAPIResponse() *TmallServiceSettleadjustmentModifyAPIResponse {
	return poolTmallServiceSettleadjustmentModifyAPIResponse.Get().(*TmallServiceSettleadjustmentModifyAPIResponse)
}

// ReleaseTmallServiceSettleadjustmentModifyAPIResponse 将 TmallServiceSettleadjustmentModifyAPIResponse 保存到 sync.Pool
func ReleaseTmallServiceSettleadjustmentModifyAPIResponse(v *TmallServiceSettleadjustmentModifyAPIResponse) {
	v.Reset()
	poolTmallServiceSettleadjustmentModifyAPIResponse.Put(v)
}
