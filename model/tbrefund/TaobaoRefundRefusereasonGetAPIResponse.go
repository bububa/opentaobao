package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundRefusereasonGetAPIResponse 获取拒绝原因列表 API返回值
// taobao.refund.refusereason.get
//
// 获取商家拒绝原因列表
type TaobaoRefundRefusereasonGetAPIResponse struct {
	model.CommonResponse
	TaobaoRefundRefusereasonGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRefundRefusereasonGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRefundRefusereasonGetAPIResponseModel).Reset()
}

// TaobaoRefundRefusereasonGetAPIResponseModel is 获取拒绝原因列表 成功返回结果
type TaobaoRefundRefusereasonGetAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_refusereason_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 卖家拒绝原因对象
	Reasons []Reason `json:"reasons,omitempty" xml:"reasons>reason,omitempty"`
	// 原因个数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRefundRefusereasonGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Reasons = m.Reasons[:0]
	m.TotalResults = 0
	m.HasNext = false
}

var poolTaobaoRefundRefusereasonGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRefundRefusereasonGetAPIResponse)
	},
}

// GetTaobaoRefundRefusereasonGetAPIResponse 从 sync.Pool 获取 TaobaoRefundRefusereasonGetAPIResponse
func GetTaobaoRefundRefusereasonGetAPIResponse() *TaobaoRefundRefusereasonGetAPIResponse {
	return poolTaobaoRefundRefusereasonGetAPIResponse.Get().(*TaobaoRefundRefusereasonGetAPIResponse)
}

// ReleaseTaobaoRefundRefusereasonGetAPIResponse 将 TaobaoRefundRefusereasonGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRefundRefusereasonGetAPIResponse(v *TaobaoRefundRefusereasonGetAPIResponse) {
	v.Reset()
	poolTaobaoRefundRefusereasonGetAPIResponse.Put(v)
}
