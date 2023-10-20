package txcs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTxcsFinanceBillConfirmAPIResponse 供应商账单确认 API返回值
// tmall.txcs.finance.bill.confirm
//
// 提供天猫超市外部合作商家：财务账单对账
type TmallTxcsFinanceBillConfirmAPIResponse struct {
	model.CommonResponse
	TmallTxcsFinanceBillConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTxcsFinanceBillConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTxcsFinanceBillConfirmAPIResponseModel).Reset()
}

// TmallTxcsFinanceBillConfirmAPIResponseModel is 供应商账单确认 成功返回结果
type TmallTxcsFinanceBillConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_txcs_finance_bill_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AccessBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTxcsFinanceBillConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTxcsFinanceBillConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTxcsFinanceBillConfirmAPIResponse)
	},
}

// GetTmallTxcsFinanceBillConfirmAPIResponse 从 sync.Pool 获取 TmallTxcsFinanceBillConfirmAPIResponse
func GetTmallTxcsFinanceBillConfirmAPIResponse() *TmallTxcsFinanceBillConfirmAPIResponse {
	return poolTmallTxcsFinanceBillConfirmAPIResponse.Get().(*TmallTxcsFinanceBillConfirmAPIResponse)
}

// ReleaseTmallTxcsFinanceBillConfirmAPIResponse 将 TmallTxcsFinanceBillConfirmAPIResponse 保存到 sync.Pool
func ReleaseTmallTxcsFinanceBillConfirmAPIResponse(v *TmallTxcsFinanceBillConfirmAPIResponse) {
	v.Reset()
	poolTmallTxcsFinanceBillConfirmAPIResponse.Put(v)
}
