package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单优惠账单查询 APIResponse
alibaba.wdk.trade.discount.bill.get

商家查询订单优惠账单
*/
type AlibabaWdkTradeDiscountBillGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkTradeDiscountBillGetResponse `json:"alibaba_wdk_trade_discount_bill_get_response,omitempty"` 
    AlibabaWdkTradeDiscountBillGetResponse
}

/* model for simplify = false
type AlibabaWdkTradeDiscountBillGetResponse struct {

    // 结果
    
    Result  *struct {
        OrderDiscountBillQueryResult  *OrderDiscountBillQueryResult `json:"order_discount_bill_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkTradeDiscountBillGetResponse struct {

    // 结果
    Result   *OrderDiscountBillQueryResult `json:"result,omitempty"`

}
