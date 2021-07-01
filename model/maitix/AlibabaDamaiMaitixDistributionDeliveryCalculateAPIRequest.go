package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest
计算渠道用户下单快递费 API请求
alibaba.damai.maitix.distribution.delivery.calculate

计算渠道用户下单快递费 */
type AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest struct {
	model.Params
	// 入参
	_param *OpenApiPostFeeParam
}

// New
