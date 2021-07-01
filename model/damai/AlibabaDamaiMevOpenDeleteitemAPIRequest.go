package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeleteitemAPIRequest
大麦换验平台-第三方对外开放-票品接口deleteItem API请求
alibaba.damai.mev.open.deleteitem

deleteItem */
type AlibabaDamaiMevOpenDeleteitemAPIRequest struct {
	model.Params
	// 入参deleteItemParam
	_deleteItemParam *TicketItemIdOpenParam
}

// New
