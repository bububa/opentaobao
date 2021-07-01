package gameact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDeActivityInfoGetAPIRequest
获取活动信息 API请求
taobao.de.activity.info.get

根据appKey和活动id获取活动 */
type TaobaoDeActivityInfoGetAPIRequest struct {
	model.Params
	// 事件唯一标识
	_eventKey string
}

// New
