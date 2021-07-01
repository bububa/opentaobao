package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkTradeDiscountBillGetAPIResponse
订单优惠账单查询 API返回值
alibaba.wdk.trade.discount.bill.get

商家查询订单优惠账单 */
type AlibabaWdkTradeDiscountBillGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeDiscountBillGetAPIResponseModel
}

// AlibabaWdkTradeDiscountBillGetAPIResponseModel is 订单优惠账单查询 成功返回结果
type AlibabaWdkTradeDiscountBillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_discount_bill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *OrderDiscountBillQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
