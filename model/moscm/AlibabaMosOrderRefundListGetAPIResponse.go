package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOrderRefundListGetAPIResponse 批量查询交易退货信息 API返回值
// alibaba.mos.order.refund.list.get
//
// 批量查询多个退货单的退货明细
type AlibabaMosOrderRefundListGetAPIResponse struct {
	model.CommonResponse
	AlibabaMosOrderRefundListGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosOrderRefundListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosOrderRefundListGetAPIResponseModel).Reset()
}

// AlibabaMosOrderRefundListGetAPIResponseModel is 批量查询交易退货信息 成功返回结果
type AlibabaMosOrderRefundListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_order_refund_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaMosOrderRefundListGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosOrderRefundListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosOrderRefundListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosOrderRefundListGetAPIResponse)
	},
}

// GetAlibabaMosOrderRefundListGetAPIResponse 从 sync.Pool 获取 AlibabaMosOrderRefundListGetAPIResponse
func GetAlibabaMosOrderRefundListGetAPIResponse() *AlibabaMosOrderRefundListGetAPIResponse {
	return poolAlibabaMosOrderRefundListGetAPIResponse.Get().(*AlibabaMosOrderRefundListGetAPIResponse)
}

// ReleaseAlibabaMosOrderRefundListGetAPIResponse 将 AlibabaMosOrderRefundListGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosOrderRefundListGetAPIResponse(v *AlibabaMosOrderRefundListGetAPIResponse) {
	v.Reset()
	poolAlibabaMosOrderRefundListGetAPIResponse.Put(v)
}
