package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundCreateAPIResponse 创建OpenMall退款单 API返回值
// taobao.openmall.refund.create
//
// 创建OpenMall退款单
// 如存在未完结的退款单，则返回该退款单ID
type TaobaoOpenmallRefundCreateAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallRefundCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallRefundCreateAPIResponseModel).Reset()
}

// TaobaoOpenmallRefundCreateAPIResponseModel is 创建OpenMall退款单 成功返回结果
type TaobaoOpenmallRefundCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_refund_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款单状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 退款ID
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Status = ""
	m.RefundId = 0
}

var poolTaobaoOpenmallRefundCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallRefundCreateAPIResponse)
	},
}

// GetTaobaoOpenmallRefundCreateAPIResponse 从 sync.Pool 获取 TaobaoOpenmallRefundCreateAPIResponse
func GetTaobaoOpenmallRefundCreateAPIResponse() *TaobaoOpenmallRefundCreateAPIResponse {
	return poolTaobaoOpenmallRefundCreateAPIResponse.Get().(*TaobaoOpenmallRefundCreateAPIResponse)
}

// ReleaseTaobaoOpenmallRefundCreateAPIResponse 将 TaobaoOpenmallRefundCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallRefundCreateAPIResponse(v *TaobaoOpenmallRefundCreateAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallRefundCreateAPIResponse.Put(v)
}
