package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSpecialRefundGetAPIResponse 特殊部分退纠纷单查询 API返回值
// taobao.special.refund.get
//
// 获取单笔特殊部分退的纠纷单查询
type TaobaoSpecialRefundGetAPIResponse struct {
	model.CommonResponse
	TaobaoSpecialRefundGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSpecialRefundGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSpecialRefundGetAPIResponseModel).Reset()
}

// TaobaoSpecialRefundGetAPIResponseModel is 特殊部分退纠纷单查询 成功返回结果
type TaobaoSpecialRefundGetAPIResponseModel struct {
	XMLName xml.Name `xml:"special_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款详情
	Refund *Refund `json:"refund,omitempty" xml:"refund,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSpecialRefundGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Refund = nil
}

var poolTaobaoSpecialRefundGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSpecialRefundGetAPIResponse)
	},
}

// GetTaobaoSpecialRefundGetAPIResponse 从 sync.Pool 获取 TaobaoSpecialRefundGetAPIResponse
func GetTaobaoSpecialRefundGetAPIResponse() *TaobaoSpecialRefundGetAPIResponse {
	return poolTaobaoSpecialRefundGetAPIResponse.Get().(*TaobaoSpecialRefundGetAPIResponse)
}

// ReleaseTaobaoSpecialRefundGetAPIResponse 将 TaobaoSpecialRefundGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSpecialRefundGetAPIResponse(v *TaobaoSpecialRefundGetAPIResponse) {
	v.Reset()
	poolTaobaoSpecialRefundGetAPIResponse.Put(v)
}
