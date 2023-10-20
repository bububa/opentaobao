package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundsApplyGetAPIResponse 查询买家申请的退款列表 API返回值
// taobao.refunds.apply.get
//
// 查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型
type TaobaoRefundsApplyGetAPIResponse struct {
	model.CommonResponse
	TaobaoRefundsApplyGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRefundsApplyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRefundsApplyGetAPIResponseModel).Reset()
}

// TaobaoRefundsApplyGetAPIResponseModel is 查询买家申请的退款列表 成功返回结果
type TaobaoRefundsApplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"refunds_apply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的退款信息列表
	Refunds []Refund `json:"refunds,omitempty" xml:"refunds>refund,omitempty"`
	// 搜索到的交易信息总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRefundsApplyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Refunds = m.Refunds[:0]
	m.TotalResults = 0
}

var poolTaobaoRefundsApplyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRefundsApplyGetAPIResponse)
	},
}

// GetTaobaoRefundsApplyGetAPIResponse 从 sync.Pool 获取 TaobaoRefundsApplyGetAPIResponse
func GetTaobaoRefundsApplyGetAPIResponse() *TaobaoRefundsApplyGetAPIResponse {
	return poolTaobaoRefundsApplyGetAPIResponse.Get().(*TaobaoRefundsApplyGetAPIResponse)
}

// ReleaseTaobaoRefundsApplyGetAPIResponse 将 TaobaoRefundsApplyGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRefundsApplyGetAPIResponse(v *TaobaoRefundsApplyGetAPIResponse) {
	v.Reset()
	poolTaobaoRefundsApplyGetAPIResponse.Put(v)
}
