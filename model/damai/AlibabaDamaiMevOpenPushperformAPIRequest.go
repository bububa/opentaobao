package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushperformAPIRequest
大麦换验平台-第三方对外开放-场次接口pushPerform API请求
alibaba.damai.mev.open.pushperform

pushPerform */
type AlibabaDamaiMevOpenPushperformAPIRequest struct {
	model.Params
	// 入参pushPerformParam
	_pushPerformParam *ThirdPerformPushOpenParam
}

// New
