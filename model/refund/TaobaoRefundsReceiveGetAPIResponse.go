package refund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundsReceiveGetAPIResponse 查询卖家收到的退款列表 API返回值
// taobao.refunds.receive.get
//
// 查询卖家收到的退款列表
type TaobaoRefundsReceiveGetAPIResponse struct {
	model.CommonResponse
	TaobaoRefundsReceiveGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRefundsReceiveGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRefundsReceiveGetAPIResponseModel).Reset()
}

// TaobaoRefundsReceiveGetAPIResponseModel is 查询卖家收到的退款列表 成功返回结果
type TaobaoRefundsReceiveGetAPIResponseModel struct {
	XMLName xml.Name `xml:"refunds_receive_get_response"`
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
func (m *TaobaoRefundsReceiveGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Refunds = m.Refunds[:0]
	m.TotalResults = 0
	m.HasNext = false
}

var poolTaobaoRefundsReceiveGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRefundsReceiveGetAPIResponse)
	},
}

// GetTaobaoRefundsReceiveGetAPIResponse 从 sync.Pool 获取 TaobaoRefundsReceiveGetAPIResponse
func GetTaobaoRefundsReceiveGetAPIResponse() *TaobaoRefundsReceiveGetAPIResponse {
	return poolTaobaoRefundsReceiveGetAPIResponse.Get().(*TaobaoRefundsReceiveGetAPIResponse)
}

// ReleaseTaobaoRefundsReceiveGetAPIResponse 将 TaobaoRefundsReceiveGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRefundsReceiveGetAPIResponse(v *TaobaoRefundsReceiveGetAPIResponse) {
	v.Reset()
	poolTaobaoRefundsReceiveGetAPIResponse.Put(v)
}
