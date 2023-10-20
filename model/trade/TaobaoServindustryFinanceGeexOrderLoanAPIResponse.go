package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceGeexOrderLoanAPIResponse 即科放款信息回调api API返回值
// taobao.servindustry.finance.geex.order.loan
//
// 即科放款信息api
type TaobaoServindustryFinanceGeexOrderLoanAPIResponse struct {
	model.CommonResponse
	TaobaoServindustryFinanceGeexOrderLoanAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoServindustryFinanceGeexOrderLoanAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoServindustryFinanceGeexOrderLoanAPIResponseModel).Reset()
}

// TaobaoServindustryFinanceGeexOrderLoanAPIResponseModel is 即科放款信息回调api 成功返回结果
type TaobaoServindustryFinanceGeexOrderLoanAPIResponseModel struct {
	XMLName xml.Name `xml:"servindustry_finance_geex_order_loan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *RetryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoServindustryFinanceGeexOrderLoanAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoServindustryFinanceGeexOrderLoanAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoServindustryFinanceGeexOrderLoanAPIResponse)
	},
}

// GetTaobaoServindustryFinanceGeexOrderLoanAPIResponse 从 sync.Pool 获取 TaobaoServindustryFinanceGeexOrderLoanAPIResponse
func GetTaobaoServindustryFinanceGeexOrderLoanAPIResponse() *TaobaoServindustryFinanceGeexOrderLoanAPIResponse {
	return poolTaobaoServindustryFinanceGeexOrderLoanAPIResponse.Get().(*TaobaoServindustryFinanceGeexOrderLoanAPIResponse)
}

// ReleaseTaobaoServindustryFinanceGeexOrderLoanAPIResponse 将 TaobaoServindustryFinanceGeexOrderLoanAPIResponse 保存到 sync.Pool
func ReleaseTaobaoServindustryFinanceGeexOrderLoanAPIResponse(v *TaobaoServindustryFinanceGeexOrderLoanAPIResponse) {
	v.Reset()
	poolTaobaoServindustryFinanceGeexOrderLoanAPIResponse.Put(v)
}
