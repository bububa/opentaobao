package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单优惠账单查询 APIResponse
alibaba.wdk.trade.discount.bill.get

商家查询订单优惠账单
*/
type AlibabaWdkTradeDiscountBillGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkTradeDiscountBillGetResponse
}

type AlibabaWdkTradeDiscountBillGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_trade_discount_bill_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *OrderDiscountBillQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
