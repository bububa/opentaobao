package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterEntityDeleteAPIRequest
删除外部成本中心人员信息 API请求
alitrip.btrip.cost.center.entity.delete

删除外部成本中心人员信息 */
type AlitripBtripCostCenterEntityDeleteAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterDeleteEntityRq
}

// New
