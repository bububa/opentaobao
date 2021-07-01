package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtipCostCenterQueryAPIRequest
查询外部成本中心 API请求
alitrip.btip.cost.center.query

查询外部成本中心 */
type AlitripBtipCostCenterQueryAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterQueryRq
}

// New
