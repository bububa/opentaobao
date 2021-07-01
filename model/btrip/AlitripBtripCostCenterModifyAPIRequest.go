package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterModifyAPIRequest
修改外部成本中心 API请求
alitrip.btrip.cost.center.modify

修改外部成本中心，设置成员，设置支付宝账号，设置名称，编号等 */
type AlitripBtripCostCenterModifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterModifyRq
}

// New
