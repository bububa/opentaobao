package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenCostCenterQueryAPIRequest
查询成本中心 API请求
alitrip.btrip.open.cost.center.query

查询成本中心 */
type AlitripBtripOpenCostCenterQueryAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterQueryRq
}

// New
