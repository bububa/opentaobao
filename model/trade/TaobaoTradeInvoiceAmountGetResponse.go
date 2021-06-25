package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取订单应开票金额 APIResponse
taobao.trade.invoice.amount.get

订单应开票金额计算
*/
type TaobaoTradeInvoiceAmountGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTradeInvoiceAmountGetResponse `json:"taobao_trade_invoice_amount_get_response,omitempty"`
}

type TaobaoTradeInvoiceAmountGetResponse struct {

    // 应开票给消费者的金额，单位分
    ConsumerInvoiceAmount   string `json:"consumer_invoice_amount,omitempty"`

    // 应开票给平台的金额，单位分
    PlatformInvoiceAmount   string `json:"platform_invoice_amount,omitempty"`

}
