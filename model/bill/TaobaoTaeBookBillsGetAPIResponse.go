package bill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaeBookBillsGetAPIResponse tae查询虚拟账户明细数据 API返回值
// taobao.tae.book.bills.get
//
// tae查询虚拟账户明细数据
type TaobaoTaeBookBillsGetAPIResponse struct {
	model.CommonResponse
	TaobaoTaeBookBillsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTaeBookBillsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTaeBookBillsGetAPIResponseModel).Reset()
}

// TaobaoTaeBookBillsGetAPIResponseModel is tae查询虚拟账户明细数据 成功返回结果
type TaobaoTaeBookBillsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tae_book_bills_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 虚拟账户账单列表
	Bills []TopAcctCashJourDto `json:"bills,omitempty" xml:"bills>top_acct_cash_jour_dto,omitempty"`
	// 当前查询的结果数,非总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTaeBookBillsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Bills = m.Bills[:0]
	m.TotalResults = 0
	m.HasNext = false
}

var poolTaobaoTaeBookBillsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTaeBookBillsGetAPIResponse)
	},
}

// GetTaobaoTaeBookBillsGetAPIResponse 从 sync.Pool 获取 TaobaoTaeBookBillsGetAPIResponse
func GetTaobaoTaeBookBillsGetAPIResponse() *TaobaoTaeBookBillsGetAPIResponse {
	return poolTaobaoTaeBookBillsGetAPIResponse.Get().(*TaobaoTaeBookBillsGetAPIResponse)
}

// ReleaseTaobaoTaeBookBillsGetAPIResponse 将 TaobaoTaeBookBillsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTaeBookBillsGetAPIResponse(v *TaobaoTaeBookBillsGetAPIResponse) {
	v.Reset()
	poolTaobaoTaeBookBillsGetAPIResponse.Put(v)
}
