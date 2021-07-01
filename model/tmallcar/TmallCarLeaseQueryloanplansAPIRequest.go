package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseQueryloanplansAPIRequest
天猫开新车租后查询还款计划 API请求
tmall.car.lease.queryloanplans

天猫开新车租后查询还款计划 */
type TmallCarLeaseQueryloanplansAPIRequest struct {
	model.Params
	// 合约编号
	_loanarno string
	// 客户的角色编号
	_iproleid string
}

// New
