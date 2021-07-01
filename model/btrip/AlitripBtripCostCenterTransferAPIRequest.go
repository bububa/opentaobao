package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterTransferAPIRequest
商旅成本中心转换为外部成本中心 API请求
alitrip.btrip.cost.center.transfer

商旅成本中心转换为外部成本中心 */
type AlitripBtripCostCenterTransferAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterTransferRq
}

// New
