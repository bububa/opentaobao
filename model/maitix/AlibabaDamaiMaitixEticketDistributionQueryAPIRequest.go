package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixEticketDistributionQueryAPIRequest
分销电子票查询接口 API请求
alibaba.damai.maitix.eticket.distribution.query

分销电子票查询接口 */
type AlibabaDamaiMaitixEticketDistributionQueryAPIRequest struct {
	model.Params
	// 入参param
	_param *EticketQueryParam
}

// New
