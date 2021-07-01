package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenCostCenterEntitySetAPIRequest
设置成本中心人员信息 API请求
alitrip.btrip.open.cost.center.entity.set

设置成本中心人员信息 */
type AlitripBtripOpenCostCenterEntitySetAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterSetEntityRq
}

// New
