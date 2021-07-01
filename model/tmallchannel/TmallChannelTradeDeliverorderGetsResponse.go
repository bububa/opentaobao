package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发货单列表 API返回值 
tmall.channel.trade.deliverorder.gets

查询发货单列表
*/
type TmallChannelTradeDeliverorderGetsAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeDeliverorderGetsResponse
}

// 查询发货单列表 成功返回结果
type TmallChannelTradeDeliverorderGetsResponse struct {
    XMLName xml.Name `xml:"tmall_channel_trade_deliverorder_gets_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 每页显示数量
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 查询第几页
    PageNumber   int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
    // 此页中包含的元素
    PageElements   []TopChannelDeliverOrderDto `json:"page_elements,omitempty" xml:"page_elements>top_channel_deliver_order_dto,omitempty"`
    // 所有元素个数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
