package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterDeleteAPIRequest
删除外部成本中心 API请求
alitrip.btrip.cost.center.delete

删除外部成本中心 */
type AlitripBtripCostCenterDeleteAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterDeleteRq
}

// New
