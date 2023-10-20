package perfect

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcwmsOutboundOrderCancelAPIResponse 取消出库单 API返回值
// alibaba.tcwms.outbound.order.cancel
//
// 取消出库单
type AlibabaTcwmsOutboundOrderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaTcwmsOutboundOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTcwmsOutboundOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTcwmsOutboundOrderCancelAPIResponseModel).Reset()
}

// AlibabaTcwmsOutboundOrderCancelAPIResponseModel is 取消出库单 成功返回结果
type AlibabaTcwmsOutboundOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcwms_outbound_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *OutboundOrderCancelResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTcwmsOutboundOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTcwmsOutboundOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTcwmsOutboundOrderCancelAPIResponse)
	},
}

// GetAlibabaTcwmsOutboundOrderCancelAPIResponse 从 sync.Pool 获取 AlibabaTcwmsOutboundOrderCancelAPIResponse
func GetAlibabaTcwmsOutboundOrderCancelAPIResponse() *AlibabaTcwmsOutboundOrderCancelAPIResponse {
	return poolAlibabaTcwmsOutboundOrderCancelAPIResponse.Get().(*AlibabaTcwmsOutboundOrderCancelAPIResponse)
}

// ReleaseAlibabaTcwmsOutboundOrderCancelAPIResponse 将 AlibabaTcwmsOutboundOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTcwmsOutboundOrderCancelAPIResponse(v *AlibabaTcwmsOutboundOrderCancelAPIResponse) {
	v.Reset()
	poolAlibabaTcwmsOutboundOrderCancelAPIResponse.Put(v)
}
