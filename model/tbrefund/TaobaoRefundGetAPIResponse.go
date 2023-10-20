package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundGetAPIResponse 获取单笔退款详情 API返回值
// taobao.refund.get
//
// 获取单笔退款详情
type TaobaoRefundGetAPIResponse struct {
	model.CommonResponse
	TaobaoRefundGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRefundGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRefundGetAPIResponseModel).Reset()
}

// TaobaoRefundGetAPIResponseModel is 获取单笔退款详情 成功返回结果
type TaobaoRefundGetAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款详情
	Refund *Refund `json:"refund,omitempty" xml:"refund,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRefundGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Refund = nil
}

var poolTaobaoRefundGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRefundGetAPIResponse)
	},
}

// GetTaobaoRefundGetAPIResponse 从 sync.Pool 获取 TaobaoRefundGetAPIResponse
func GetTaobaoRefundGetAPIResponse() *TaobaoRefundGetAPIResponse {
	return poolTaobaoRefundGetAPIResponse.Get().(*TaobaoRefundGetAPIResponse)
}

// ReleaseTaobaoRefundGetAPIResponse 将 TaobaoRefundGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRefundGetAPIResponse(v *TaobaoRefundGetAPIResponse) {
	v.Reset()
	poolTaobaoRefundGetAPIResponse.Put(v)
}
