package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenCostCenterEntityDeleteAPIRequest
删除成本中心人员信息 API请求
alitrip.btrip.open.cost.center.entity.delete

删除成本中心人员信息 */
type AlitripBtripOpenCostCenterEntityDeleteAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterDeleteEntityRq
}

// New
