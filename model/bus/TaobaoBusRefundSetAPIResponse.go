package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusRefundSetAPIResponse B2B退票申请接口 API返回值
// taobao.bus.refund.set
//
// B2B业务支持退票
type TaobaoBusRefundSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusRefundSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusRefundSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusRefundSetAPIResponseModel).Reset()
}

// TaobaoBusRefundSetAPIResponseModel is B2B退票申请接口 成功返回结果
type TaobaoBusRefundSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_refund_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *B2BRefundOrderRp `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusRefundSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBusRefundSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusRefundSetAPIResponse)
	},
}

// GetTaobaoBusRefundSetAPIResponse 从 sync.Pool 获取 TaobaoBusRefundSetAPIResponse
func GetTaobaoBusRefundSetAPIResponse() *TaobaoBusRefundSetAPIResponse {
	return poolTaobaoBusRefundSetAPIResponse.Get().(*TaobaoBusRefundSetAPIResponse)
}

// ReleaseTaobaoBusRefundSetAPIResponse 将 TaobaoBusRefundSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusRefundSetAPIResponse(v *TaobaoBusRefundSetAPIResponse) {
	v.Reset()
	poolTaobaoBusRefundSetAPIResponse.Put(v)
}
