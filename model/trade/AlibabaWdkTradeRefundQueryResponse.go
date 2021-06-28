package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道查询退货订单详情接口 APIResponse
alibaba.wdk.trade.refund.query

该接口提供给外部渠道商家，比如欧尚外卖等查询退货订单详情，里面包含退货进度等信息。
*/
type AlibabaWdkTradeRefundQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkTradeRefundQueryResponse `json:"alibaba_wdk_trade_refund_query_response,omitempty"` 
    AlibabaWdkTradeRefundQueryResponse
}

/* model for simplify = false
type AlibabaWdkTradeRefundQueryResponse struct {

    // 查询结果
    
    RefundGoodsQueryResult  *struct {
        RefundGoodsQueryResult  *RefundGoodsQueryResult `json:"refund_goods_query_result,omitempty"`
    } `json:"refund_goods_query_result,omitempty"`
    

}
*/

type AlibabaWdkTradeRefundQueryResponse struct {

    // 查询结果
    RefundGoodsQueryResult   *RefundGoodsQueryResult `json:"refund_goods_query_result,omitempty"`

}
