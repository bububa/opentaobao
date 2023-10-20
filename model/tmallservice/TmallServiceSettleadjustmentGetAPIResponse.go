package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentGetAPIResponse 查询结算调整单单条记录 API返回值
// tmall.service.settleadjustment.get
//
// 提供给服务商通过结算调整单id获取结算调整单信息
type TmallServiceSettleadjustmentGetAPIResponse struct {
	model.CommonResponse
	TmallServiceSettleadjustmentGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServiceSettleadjustmentGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServiceSettleadjustmentGetAPIResponseModel).Reset()
}

// TmallServiceSettleadjustmentGetAPIResponseModel is 查询结算调整单单条记录 成功返回结果
type TmallServiceSettleadjustmentGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settleadjustment_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServiceSettleadjustmentGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServiceSettleadjustmentGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServiceSettleadjustmentGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentGetAPIResponse)
	},
}

// GetTmallServiceSettleadjustmentGetAPIResponse 从 sync.Pool 获取 TmallServiceSettleadjustmentGetAPIResponse
func GetTmallServiceSettleadjustmentGetAPIResponse() *TmallServiceSettleadjustmentGetAPIResponse {
	return poolTmallServiceSettleadjustmentGetAPIResponse.Get().(*TmallServiceSettleadjustmentGetAPIResponse)
}

// ReleaseTmallServiceSettleadjustmentGetAPIResponse 将 TmallServiceSettleadjustmentGetAPIResponse 保存到 sync.Pool
func ReleaseTmallServiceSettleadjustmentGetAPIResponse(v *TmallServiceSettleadjustmentGetAPIResponse) {
	v.Reset()
	poolTmallServiceSettleadjustmentGetAPIResponse.Put(v)
}
