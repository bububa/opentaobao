package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushstandAPIRequest
大麦换验平台-第三方对外开放-看台接口pushStand API请求
alibaba.damai.mev.open.pushstand

pushStand */
type AlibabaDamaiMevOpenPushstandAPIRequest struct {
	model.Params
	// 入参pushStandParam
	_pushStandParam *ThirdStandPushOpenParam
}

// New
