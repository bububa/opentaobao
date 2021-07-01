package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushitemAPIRequest
大麦换验平台-第三方对外开放-票品接口pushItem API请求
alibaba.damai.mev.open.pushitem

开放接口 推送票品 */
type AlibabaDamaiMevOpenPushitemAPIRequest struct {
	model.Params
	// 入参pushItemParam
	_pushItemParam *PushTicketItemPushOpenParam
}

// New
