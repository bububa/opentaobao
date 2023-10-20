package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse 10-同步库存现有量 API返回值
// alibaba.tmallgenie.scp.plan.inventor.qty.get
//
// 同步库存现有量
type AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanInventorQtyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanInventorQtyGetAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanInventorQtyGetAPIResponseModel is 10-同步库存现有量 成功返回结果
type AlibabaTmallgenieScpPlanInventorQtyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_inventor_qty_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanInventorQtyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTmallgenieScpPlanInventorQtyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanInventorQtyGetAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse
func GetAlibabaTmallgenieScpPlanInventorQtyGetAPIResponse() *AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse {
	return poolAlibabaTmallgenieScpPlanInventorQtyGetAPIResponse.Get().(*AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanInventorQtyGetAPIResponse 将 AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanInventorQtyGetAPIResponse(v *AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanInventorQtyGetAPIResponse.Put(v)
}
