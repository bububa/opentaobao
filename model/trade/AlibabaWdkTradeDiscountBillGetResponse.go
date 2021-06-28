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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_trade_discount_bill_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *OrderDiscountBillQueryResult `json:"result,omitempty" xml:"