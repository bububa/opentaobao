package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest
分销状态查询接口queryPerformStatusByPerformId API请求
alibaba.damai.maitix.opengateway.perform.status.query

queryPerformStatusByPerformId */
type AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest struct {
	model.Params
	// 入参
	_disPerformStatusQueryParam *DisPerformStatusQueryDto
}

// New
