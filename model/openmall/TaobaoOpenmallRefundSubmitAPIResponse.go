package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundSubmitAPIResponse 提交OpenMall退款单物流 API返回值
// taobao.openmall.refund.submit
//
// 提交OpenMall退款单物流
type TaobaoOpenmallRefundSubmitAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallRefundSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallRefundSubmitAPIResponseModel).Reset()
}

// TaobaoOpenmallRefundSubmitAPIResponseModel is 提交OpenMall退款单物流 成功返回结果
type TaobaoOpenmallRefundSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_refund_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 提交物流单成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoOpenmallRefundSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallRefundSubmitAPIResponse)
	},
}

// GetTaobaoOpenmallRefundSubmitAPIResponse 从 sync.Pool 获取 TaobaoOpenmallRefundSubmitAPIResponse
func GetTaobaoOpenmallRefundSubmitAPIResponse() *TaobaoOpenmallRefundSubmitAPIResponse {
	return poolTaobaoOpenmallRefundSubmitAPIResponse.Get().(*TaobaoOpenmallRefundSubmitAPIResponse)
}

// ReleaseTaobaoOpenmallRefundSubmitAPIResponse 将 TaobaoOpenmallRefundSubmitAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallRefundSubmitAPIResponse(v *TaobaoOpenmallRefundSubmitAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallRefundSubmitAPIResponse.Put(v)
}
