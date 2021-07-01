package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeletefloorAPIRequest
大麦换验平台-第三方对外开放-楼层接口deleteFloor API请求
alibaba.damai.mev.open.deletefloor

deleteFloor */
type AlibabaDamaiMevOpenDeletefloorAPIRequest struct {
	model.Params
	// 入参deleteFloorParam
	_deleteFloorParam *FloorIdOpenParam
}

// New
