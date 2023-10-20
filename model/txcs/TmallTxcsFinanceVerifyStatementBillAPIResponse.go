package txcs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTxcsFinanceVerifyStatementBillAPIResponse 供应商核销单录入 API返回值
// tmall.txcs.finance.verify.statement.bill
//
// 供应商核销单录入
type TmallTxcsFinanceVerifyStatementBillAPIResponse struct {
	model.CommonResponse
	TmallTxcsFinanceVerifyStatementBillAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTxcsFinanceVerifyStatementBillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTxcsFinanceVerifyStatementBillAPIResponseModel).Reset()
}

// TmallTxcsFinanceVerifyStatementBillAPIResponseModel is 供应商核销单录入 成功返回结果
type TmallTxcsFinanceVerifyStatementBillAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_txcs_finance_verify_statement_bill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AccessBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTxcsFinanceVerifyStatementBillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTxcsFinanceVerifyStatementBillAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTxcsFinanceVerifyStatementBillAPIResponse)
	},
}

// GetTmallTxcsFinanceVerifyStatementBillAPIResponse 从 sync.Pool 获取 TmallTxcsFinanceVerifyStatementBillAPIResponse
func GetTmallTxcsFinanceVerifyStatementBillAPIResponse() *TmallTxcsFinanceVerifyStatementBillAPIResponse {
	return poolTmallTxcsFinanceVerifyStatementBillAPIResponse.Get().(*TmallTxcsFinanceVerifyStatementBillAPIResponse)
}

// ReleaseTmallTxcsFinanceVerifyStatementBillAPIResponse 将 TmallTxcsFinanceVerifyStatementBillAPIResponse 保存到 sync.Pool
func ReleaseTmallTxcsFinanceVerifyStatementBillAPIResponse(v *TmallTxcsFinanceVerifyStatementBillAPIResponse) {
	v.Reset()
	poolTmallTxcsFinanceVerifyStatementBillAPIResponse.Put(v)
}
