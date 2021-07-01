package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenCostCenterEntityAddAPIRequest
增加成本中心人员信息 API请求
alitrip.btrip.open.cost.center.entity.add

增加成本中心人员信息 */
type AlitripBtripOpenCostCenterEntityAddAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterAddEntityRq
}

// New
