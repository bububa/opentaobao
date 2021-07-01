package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeInvoiceAmountGetAPIResponse
获取订单应开票金额 API返回值
taobao.trade.invoice.amount.get

订单应开票金额计算 */
type TaobaoTradeInvoiceAmountGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeInvoiceAmountGetAPIResponseModel
}

// TaobaoTradeInvoiceAmountGetAPIResponseModel is 获取订单应开票金额 成功返回结果
type TaobaoTradeInvoiceAmountGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_invoice_amount_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 应开票给消费者的金额，单位分
	ConsumerInvoiceAmount string `json:"consumer_invoice_amount,omitempty" xml:"consumer_invoice_amount,omitempty"`
	// 应开票给平台的金额，单位分
	PlatformInvoiceAmount string `json:"platform_invoice_amount,omitempty" xml:"platform_invoice_amount,omitempty"`
}
