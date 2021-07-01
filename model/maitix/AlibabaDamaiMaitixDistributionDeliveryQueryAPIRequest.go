package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest
查询分销物流单 API请求
alibaba.damai.maitix.distribution.delivery.query

渠道查询物流订单 */
type AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest struct {
	model.Params
	// 主订单号
	_mainOrderId string
}

// New
