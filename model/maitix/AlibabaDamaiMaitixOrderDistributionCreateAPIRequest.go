package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixOrderDistributionCreateAPIRequest
大麦-新分销下单 API请求
alibaba.damai.maitix.order.distribution.create

createDistributionOrder */
type AlibabaDamaiMaitixOrderDistributionCreateAPIRequest struct {
	model.Params
	// 下单参数param
	_param *MoaOrderParam
}

// New
