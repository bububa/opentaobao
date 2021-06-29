package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取采购申请单列表 APIResponse
tmall.channel.trade.applyorder.gets

分页查询采购申请单列表
*/
type TmallChannelTradeApplyorderGetsAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeApplyorderGetsResponse
}

type TmallChannelTradeApplyorderGetsResponse struct {
    XMLName xml.Name `xml:"tmall_channel_trade_applyorder_gets_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 每页显示数量
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`

    
    // 查询第几页
    
    PageNumber   int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`

    
    // 包含的元素
    
    PageElements   []TopChannelApplyOrderDto `json:"page_elements,omitempty" xml:"page_elements>top_channel_apply_order_dto,omitempty"`
    
    
    // 所有元素个数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
}
