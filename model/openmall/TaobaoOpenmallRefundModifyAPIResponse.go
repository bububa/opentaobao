package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundModifyAPIResponse 修改OpenMall退款申请 API返回值
// taobao.openmall.refund.modify
//
// 修改OpenMall退款申请
type TaobaoOpenmallRefundModifyAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallRefundModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallRefundModifyAPIResponseModel).Reset()
}

// TaobaoOpenmallRefundModifyAPIResponseModel is 修改OpenMall退款申请 成功返回结果
type TaobaoOpenmallRefundModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_refund_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoOpenmallRefundModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallRefundModifyAPIResponse)
	},
}

// GetTaobaoOpenmallRefundModifyAPIResponse 从 sync.Pool 获取 TaobaoOpenmallRefundModifyAPIResponse
func GetTaobaoOpenmallRefundModifyAPIResponse() *TaobaoOpenmallRefundModifyAPIResponse {
	return poolTaobaoOpenmallRefundModifyAPIResponse.Get().(*TaobaoOpenmallRefundModifyAPIResponse)
}

// ReleaseTaobaoOpenmallRefundModifyAPIResponse 将 TaobaoOpenmallRefundModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallRefundModifyAPIResponse(v *TaobaoOpenmallRefundModifyAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallRefundModifyAPIResponse.Put(v)
}
