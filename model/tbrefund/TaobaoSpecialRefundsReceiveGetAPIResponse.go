package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSpecialRefundsReceiveGetAPIResponse 特殊退款类型的纠纷单列表查询 API返回值
// taobao.special.refunds.receive.get
//
// 特殊退款类型的纠纷单列表查询
type TaobaoSpecialRefundsReceiveGetAPIResponse struct {
	model.CommonResponse
	TaobaoSpecialRefundsReceiveGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSpecialRefundsReceiveGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSpecialRefundsReceiveGetAPIResponseModel).Reset()
}

// TaobaoSpecialRefundsReceiveGetAPIResponseModel is 特殊退款类型的纠纷单列表查询 成功返回结果
type TaobaoSpecialRefundsReceiveGetAPIResponseModel struct {
	XMLName xml.Name `xml:"special_refunds_receive_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的退款信息列表
	Refunds []Refund `json:"refunds,omitempty" xml:"refunds>refund,omitempty"`
	// 搜索到的退款信息总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSpecialRefundsReceiveGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Refunds = m.Refunds[:0]
	m.TotalResults = 0
	m.HasNext = false
}

var poolTaobaoSpecialRefundsReceiveGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSpecialRefundsReceiveGetAPIResponse)
	},
}

// GetTaobaoSpecialRefundsReceiveGetAPIResponse 从 sync.Pool 获取 TaobaoSpecialRefundsReceiveGetAPIResponse
func GetTaobaoSpecialRefundsReceiveGetAPIResponse() *TaobaoSpecialRefundsReceiveGetAPIResponse {
	return poolTaobaoSpecialRefundsReceiveGetAPIResponse.Get().(*TaobaoSpecialRefundsReceiveGetAPIResponse)
}

// ReleaseTaobaoSpecialRefundsReceiveGetAPIResponse 将 TaobaoSpecialRefundsReceiveGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSpecialRefundsReceiveGetAPIResponse(v *TaobaoSpecialRefundsReceiveGetAPIResponse) {
	v.Reset()
	poolTaobaoSpecialRefundsReceiveGetAPIResponse.Put(v)
}
