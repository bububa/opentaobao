package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcQueueGetAPIRequest
获取消息队列积压情况 API请求
taobao.tmc.queue.get

根据appkey和groupName获取消息队列积压情况 */
type TaobaoTmcQueueGetAPIRequest struct {
	model.Params
	// TMC组名
	_groupName string
}

// New
