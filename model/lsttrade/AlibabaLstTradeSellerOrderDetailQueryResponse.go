package lsttrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详情查看(卖家视角) APIResponse
alibaba.lst.trade.seller.order.detail.query

订单详情查看(卖家视角)
*/
type AlibabaLstTradeSellerOrderDetailQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstTradeSellerOrderDetailQueryResponse
}

type AlibabaLstTradeSellerOrderDetailQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_trade_seller_order_detail_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlibabaLstTradeSellerOrderDetailQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
