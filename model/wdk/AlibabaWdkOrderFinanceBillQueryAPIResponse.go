package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOrderFinanceBillQueryAPIResponse
资金合规商家账单 API返回值
alibaba.wdk.order.finance.bill.query

拉取资金合规商家账单 */
type AlibabaWdkOrderFinanceBillQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOrderFinanceBillQueryAPIResponseModel
}

// AlibabaWdkOrderFinanceBillQueryAPIResponseModel is 资金合规商家账单 成功返回结果
type AlibabaWdkOrderFinanceBillQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_finance_bill_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *WdkOpenOrderFinanceBillQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
