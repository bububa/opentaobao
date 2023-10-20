package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse 12-同步销售数据 API返回值
// alibaba.tmallgenie.scp.plan.sale.qty.get
//
// 同步销售数据
type AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanSaleQtyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanSaleQtyGetAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanSaleQtyGetAPIResponseModel is 12-同步销售数据 成功返回结果
type AlibabaTmallgenieScpPlanSaleQtyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_sale_qty_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanSaleQtyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTmallgenieScpPlanSaleQtyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanSaleQtyGetAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse
func GetAlibabaTmallgenieScpPlanSaleQtyGetAPIResponse() *AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse {
	return poolAlibabaTmallgenieScpPlanSaleQtyGetAPIResponse.Get().(*AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanSaleQtyGetAPIResponse 将 AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanSaleQtyGetAPIResponse(v *AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanSaleQtyGetAPIResponse.Put(v)
}
