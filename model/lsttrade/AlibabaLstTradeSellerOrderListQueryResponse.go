package lsttrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单列表查看(卖家视角) APIResponse
alibaba.lst.trade.seller.order.list.query

卖家视角订单查询，查询授权经销商订单列表
*/
type AlibabaLstTradeSellerOrderListQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstTradeSellerOrderListQueryResponse
}

type AlibabaLstTradeSellerOrderListQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_trade_seller_order_list_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaLstTradeSellerOrderListQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
