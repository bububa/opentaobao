package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusRefundticketpriceSetAPIResponse 汽车票退款申请接口 API返回值
// taobao.bus.refundticketprice.set
//
// 汽车票代理商利用该接口申请退票
type TaobaoBusRefundticketpriceSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusRefundticketpriceSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusRefundticketpriceSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusRefundticketpriceSetAPIResponseModel).Reset()
}

// TaobaoBusRefundticketpriceSetAPIResponseModel is 汽车票退款申请接口 成功返回结果
type TaobaoBusRefundticketpriceSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_refundticketprice_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退票成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusRefundticketpriceSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoBusRefundticketpriceSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusRefundticketpriceSetAPIResponse)
	},
}

// GetTaobaoBusRefundticketpriceSetAPIResponse 从 sync.Pool 获取 TaobaoBusRefundticketpriceSetAPIResponse
func GetTaobaoBusRefundticketpriceSetAPIResponse() *TaobaoBusRefundticketpriceSetAPIResponse {
	return poolTaobaoBusRefundticketpriceSetAPIResponse.Get().(*TaobaoBusRefundticketpriceSetAPIResponse)
}

// ReleaseTaobaoBusRefundticketpriceSetAPIResponse 将 TaobaoBusRefundticketpriceSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusRefundticketpriceSetAPIResponse(v *TaobaoBusRefundticketpriceSetAPIResponse) {
	v.Reset()
	poolTaobaoBusRefundticketpriceSetAPIResponse.Put(v)
}
