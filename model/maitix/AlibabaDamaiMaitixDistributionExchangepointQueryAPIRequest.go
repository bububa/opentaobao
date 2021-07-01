package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest
分销查询取票点接口 API请求
alibaba.damai.maitix.distribution.exchangepoint.query

分销查询取票点接口 */
type AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest struct {
	model.Params
	// 必填-项目id
	_projectId int64
}

// New
