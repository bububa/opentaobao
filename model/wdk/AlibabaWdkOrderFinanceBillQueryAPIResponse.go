package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderFinanceBillQueryAPIResponse 资金合规商家账单 API返回值
// alibaba.wdk.order.finance.bill.query
//
// 拉取资金合规商家账单
type AlibabaWdkOrderFinanceBillQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOrderFinanceBillQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkOrderFinanceBillQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkOrderFinanceBillQueryAPIResponseModel).Reset()
}

// AlibabaWdkOrderFinanceBillQueryAPIResponseModel is 资金合规商家账单 成功返回结果
type AlibabaWdkOrderFinanceBillQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_finance_bill_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *WdkOpenOrderFinanceBillQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkOrderFinanceBillQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkOrderFinanceBillQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkOrderFinanceBillQueryAPIResponse)
	},
}

// GetAlibabaWdkOrderFinanceBillQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkOrderFinanceBillQueryAPIResponse
func GetAlibabaWdkOrderFinanceBillQueryAPIResponse() *AlibabaWdkOrderFinanceBillQueryAPIResponse {
	return poolAlibabaWdkOrderFinanceBillQueryAPIResponse.Get().(*AlibabaWdkOrderFinanceBillQueryAPIResponse)
}

// ReleaseAlibabaWdkOrderFinanceBillQueryAPIResponse 将 AlibabaWdkOrderFinanceBillQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkOrderFinanceBillQueryAPIResponse(v *AlibabaWdkOrderFinanceBillQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkOrderFinanceBillQueryAPIResponse.Put(v)
}
