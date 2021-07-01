package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenEventsProduceAPIRequest
批量发送奇门事件 API请求
taobao.qimen.events.produce

批量发送消息 */
type TaobaoQimenEventsProduceAPIRequest struct {
	model.Params
	// 奇门事件列表, 最多50条
	_messages []QimenEvent
}

// New
