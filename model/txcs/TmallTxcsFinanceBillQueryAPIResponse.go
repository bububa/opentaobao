package txcs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTxcsFinanceBillQueryAPIResponse 天猫超市外部商家财务账单信息查询 API返回值
// tmall.txcs.finance.bill.query
//
// 提供天猫超市外部合作商家财务账单对账
type TmallTxcsFinanceBillQueryAPIResponse struct {
	model.CommonResponse
	TmallTxcsFinanceBillQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTxcsFinanceBillQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTxcsFinanceBillQueryAPIResponseModel).Reset()
}

// TmallTxcsFinanceBillQueryAPIResponseModel is 天猫超市外部商家财务账单信息查询 成功返回结果
type TmallTxcsFinanceBillQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_txcs_finance_bill_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTxcsFinanceBillQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTxcsFinanceBillQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTxcsFinanceBillQueryAPIResponse)
	},
}

// GetTmallTxcsFinanceBillQueryAPIResponse 从 sync.Pool 获取 TmallTxcsFinanceBillQueryAPIResponse
func GetTmallTxcsFinanceBillQueryAPIResponse() *TmallTxcsFinanceBillQueryAPIResponse {
	return poolTmallTxcsFinanceBillQueryAPIResponse.Get().(*TmallTxcsFinanceBillQueryAPIResponse)
}

// ReleaseTmallTxcsFinanceBillQueryAPIResponse 将 TmallTxcsFinanceBillQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallTxcsFinanceBillQueryAPIResponse(v *TmallTxcsFinanceBillQueryAPIResponse) {
	v.Reset()
	poolTmallTxcsFinanceBillQueryAPIResponse.Put(v)
}
