package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建渠道分销单 APIResponse
tmall.channel.trade.order.create

创建渠道分销单
*/
type TmallChannelTradeOrderCreateAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeOrderCreateResponse
}

type TmallChannelTradeOrderCreateResponse struct {
    XMLName xml.Name `xml:"tmall_channel_trade_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 采购单号
    
    MainPurchaseOrderList   []string `json:"main_purchase_order_list,omitempty" xml:"main_purchase_order_list>string,omitempty"`
    
    
}
