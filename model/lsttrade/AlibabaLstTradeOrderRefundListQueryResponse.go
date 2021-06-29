package lsttrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询退款单列表(卖家视角) APIResponse
alibaba.lst.trade.order.refund.list.query

查询退款单列表(卖家视角)
*/
type AlibabaLstTradeOrderRefundListQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstTradeOrderRefundListQueryResponse
}

type AlibabaLstTradeOrderRefundListQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_trade_order_refund_list_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 总数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`

    
    // 每页数量
    
    Size   int64 `json:"size,omitempty" xml:"size,omitempty"`

    
    // 当前页
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`

    
    // 列表数据
    
    ContentList   []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
    
    
}
