package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道查询退货订单详情接口 APIResponse
alibaba.wdk.trade.refund.query

该接口提供给外部渠道商家，比如欧尚外卖等查询退货订单详情，里面包含退货进度等信息。
*/
type AlibabaWdkTradeRefundQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkTradeRefundQueryResponse
}

type AlibabaWdkTradeRefundQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_trade_refund_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果
    
    RefundGoodsQueryResult   *RefundGoodsQueryResult `json:"refund_goods_query_result,omitempty" xml:"refund_goods_query_result,omitempty"`

    
}
