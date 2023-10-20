package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeDiscountBillGetAPIResponse 订单优惠账单查询 API返回值
// alibaba.wdk.trade.discount.bill.get
//
// 商家查询订单优惠账单
type AlibabaWdkTradeDiscountBillGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeDiscountBillGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTradeDiscountBillGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTradeDiscountBillGetAPIResponseModel).Reset()
}

// AlibabaWdkTradeDiscountBillGetAPIResponseModel is 订单优惠账单查询 成功返回结果
type AlibabaWdkTradeDiscountBillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_discount_bill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *OrderDiscountBillQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTradeDiscountBillGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkTradeDiscountBillGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeDiscountBillGetAPIResponse)
	},
}

// GetAlibabaWdkTradeDiscountBillGetAPIResponse 从 sync.Pool 获取 AlibabaWdkTradeDiscountBillGetAPIResponse
func GetAlibabaWdkTradeDiscountBillGetAPIResponse() *AlibabaWdkTradeDiscountBillGetAPIResponse {
	return poolAlibabaWdkTradeDiscountBillGetAPIResponse.Get().(*AlibabaWdkTradeDiscountBillGetAPIResponse)
}

// ReleaseAlibabaWdkTradeDiscountBillGetAPIResponse 将 AlibabaWdkTradeDiscountBillGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTradeDiscountBillGetAPIResponse(v *AlibabaWdkTradeDiscountBillGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkTradeDiscountBillGetAPIResponse.Put(v)
}
