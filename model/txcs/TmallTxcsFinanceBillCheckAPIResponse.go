package txcs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTxcsFinanceBillCheckAPIResponse 天猫超市外部商家财务账单对账 API返回值
// tmall.txcs.finance.bill.check
//
// 提供天猫超市外部合作商家财务账单对账
type TmallTxcsFinanceBillCheckAPIResponse struct {
	model.CommonResponse
	TmallTxcsFinanceBillCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTxcsFinanceBillCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTxcsFinanceBillCheckAPIResponseModel).Reset()
}

// TmallTxcsFinanceBillCheckAPIResponseModel is 天猫超市外部商家财务账单对账 成功返回结果
type TmallTxcsFinanceBillCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_txcs_finance_bill_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *AccessBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTxcsFinanceBillCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTxcsFinanceBillCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTxcsFinanceBillCheckAPIResponse)
	},
}

// GetTmallTxcsFinanceBillCheckAPIResponse 从 sync.Pool 获取 TmallTxcsFinanceBillCheckAPIResponse
func GetTmallTxcsFinanceBillCheckAPIResponse() *TmallTxcsFinanceBillCheckAPIResponse {
	return poolTmallTxcsFinanceBillCheckAPIResponse.Get().(*TmallTxcsFinanceBillCheckAPIResponse)
}

// ReleaseTmallTxcsFinanceBillCheckAPIResponse 将 TmallTxcsFinanceBillCheckAPIResponse 保存到 sync.Pool
func ReleaseTmallTxcsFinanceBillCheckAPIResponse(v *TmallTxcsFinanceBillCheckAPIResponse) {
	v.Reset()
	poolTmallTxcsFinanceBillCheckAPIResponse.Put(v)
}
