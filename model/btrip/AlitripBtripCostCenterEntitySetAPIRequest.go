package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterEntitySetAPIRequest
设置外部成本中心人员信息 API请求
alitrip.btrip.cost.center.entity.set

设置外部成本中心人员信息 */
type AlitripBtripCostCenterEntitySetAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterSetEntityRq
}

// New
