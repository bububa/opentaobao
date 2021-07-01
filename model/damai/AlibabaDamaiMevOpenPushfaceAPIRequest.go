package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushfaceAPIRequest
大麦换验平台-第三方对外开放-票面接口pushFace API请求
alibaba.damai.mev.open.pushface

pushFace */
type AlibabaDamaiMevOpenPushfaceAPIRequest struct {
	model.Params
	// 入参pushFaceParam
	_pushFaceParam *ThirdTicketFacePushOpenParam
}

// New
