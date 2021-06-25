package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流宝通知消息查询接口 APIResponse
taobao.wlb.notify.message.page.get

物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
*/
type TaobaoWlbNotifyMessagePageGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbNotifyMessagePageGetResponse `json:"taobao_wlb_notify_message_page_get_response,omitempty"`
}

type TaobaoWlbNotifyMessagePageGetResponse struct {

    // 2000-01-01 00:00:00
    TotalCount   int64 `json:"total_count,omitempty"`

    // 通道消息
    WlbMessages   []WlbMessage `json:"wlb_messages,omitempty"`

}
