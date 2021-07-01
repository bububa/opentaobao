package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeletefaceAPIRequest
大麦换验平台-第三方对外开放-票面接口deleteFace API请求
alibaba.damai.mev.open.deleteface

deleteFace */
type AlibabaDamaiMevOpenDeletefaceAPIRequest struct {
	model.Params
	// 入参deleteFaceParam
	_deleteFaceParam *TicketFaceIdOpenParam
}

// New
