package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseOrderidGetAPIRequest
天猫开新车查询订单id API请求
tmall.car.lease.orderid.get

天猫开新车查询订单id */
type TmallCarLeaseOrderidGetAPIRequest struct {
	model.Params
	// openid
	_openId string
}

// New
