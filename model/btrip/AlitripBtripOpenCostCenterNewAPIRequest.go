package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenCostCenterNewAPIRequest
新增成本中心 API请求
alitrip.btrip.open.cost.center.new

新增成本中心 */
type AlitripBtripOpenCostCenterNewAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterSaveRq
}

// New
