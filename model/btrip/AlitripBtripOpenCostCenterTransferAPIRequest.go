package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenCostCenterTransferAPIRequest
商旅成本中心转换为外部成本中心 API请求
alitrip.btrip.open.cost.center.transfer

商旅成本中心转换为外部成本中心 */
type AlitripBtripOpenCostCenterTransferAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterTransferRq
}

// New
