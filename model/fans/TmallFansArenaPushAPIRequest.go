package fans

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallFansArenaPushAPIRequest
消息推送 API请求
tmall.fans.arena.push

超级擂台消息推送 */
type TmallFansArenaPushAPIRequest struct {
	model.Params
	// 推送列表
	_pushList []PushMessageParamDo
}

// New
