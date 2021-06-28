package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流宝通知消息查询接口 APIResponse
taobao.wlb.notify.message.page.get

物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
*/
type TaobaoWlbNotifyMessagePageGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_notify_message_page_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 2000-01-01 00:00:00
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"