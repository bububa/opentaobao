package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbNotifyMessagePageGet 物流宝通知消息查询接口
// taobao.wlb.notify.message.page.get
//
// 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
func TaobaoWlbNotifyMessagePageGet(clt *core.SDKClient, req *wlb.TaobaoWlbNotifyMessagePageGetAPIRequest, resp *wlb.TaobaoWlbNotifyMessagePageGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
