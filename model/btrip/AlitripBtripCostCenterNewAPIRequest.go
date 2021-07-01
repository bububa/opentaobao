package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterNewAPIRequest
新建外部成本中心 API请求
alitrip.btrip.cost.center.new

新建外部成本中心 */
type AlitripBtripCostCenterNewAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterSaveRq
}

// New
