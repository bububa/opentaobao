package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenCostCenterModifyAPIRequest
修改成本中心 API请求
alitrip.btrip.open.cost.center.modify

修改成本中心 */
type AlitripBtripOpenCostCenterModifyAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterModifyRq
}

// New
