package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenCostCenterDeleteAPIRequest
删除成本中心 API请求
alitrip.btrip.open.cost.center.delete

删除成本中心 */
type AlitripBtripOpenCostCenterDeleteAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterDeleteRq
}

// New
