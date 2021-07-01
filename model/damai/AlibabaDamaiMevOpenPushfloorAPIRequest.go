package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushfloorAPIRequest
大麦换验平台-第三方对外开放-楼层接口pushFloor API请求
alibaba.damai.mev.open.pushfloor

pushFloor */
type AlibabaDamaiMevOpenPushfloorAPIRequest struct {
	model.Params
	// 入参pushFloorParam
	_pushFloorParam *ThirdFloorPushOpenParam
}

// New
