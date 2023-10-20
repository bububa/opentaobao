package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcQueueGet 获取消息队列积压情况
// taobao.tmc.queue.get
//
// 根据appkey和groupName获取消息队列积压情况
func TaobaoTmcQueueGet(clt *core.SDKClient, req *tmc.TaobaoTmcQueueGetAPIRequest, resp *tmc.TaobaoTmcQueueGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
