package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundMessageSubmitAPIResponse 提交退款单留言 API返回值
// taobao.openmall.refund.message.submit
//
// OpenMall业务提交退款单留言
type TaobaoOpenmallRefundMessageSubmitAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallRefundMessageSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundMessageSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallRefundMessageSubmitAPIResponseModel).Reset()
}

// TaobaoOpenmallRefundMessageSubmitAPIResponseModel is 提交退款单留言 成功返回结果
type TaobaoOpenmallRefundMessageSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_refund_message_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 提交结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundMessageSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoOpenmallRefundMessageSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallRefundMessageSubmitAPIResponse)
	},
}

// GetTaobaoOpenmallRefundMessageSubmitAPIResponse 从 sync.Pool 获取 TaobaoOpenmallRefundMessageSubmitAPIResponse
func GetTaobaoOpenmallRefundMessageSubmitAPIResponse() *TaobaoOpenmallRefundMessageSubmitAPIResponse {
	return poolTaobaoOpenmallRefundMessageSubmitAPIResponse.Get().(*TaobaoOpenmallRefundMessageSubmitAPIResponse)
}

// ReleaseTaobaoOpenmallRefundMessageSubmitAPIResponse 将 TaobaoOpenmallRefundMessageSubmitAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallRefundMessageSubmitAPIResponse(v *TaobaoOpenmallRefundMessageSubmitAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallRefundMessageSubmitAPIResponse.Put(v)
}
