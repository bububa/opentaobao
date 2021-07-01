package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterEntityAddAPIRequest
增加外部成本中心人员信息 API请求
alitrip.btrip.cost.center.entity.add

增加外部成本中心人员信息 */
type AlitripBtripCostCenterEntityAddAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterAddEntityRq
}

// New
