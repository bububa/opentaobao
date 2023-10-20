package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopMessageGetUnreadCount 获取未读的消息数量
// taobao.ailab.aicloud.top.message.get.unread.count
//
// 开放获取未读留言数量的接口
func TaobaoAilabAicloudTopMessageGetUnreadCount(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest, resp *iot.TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
