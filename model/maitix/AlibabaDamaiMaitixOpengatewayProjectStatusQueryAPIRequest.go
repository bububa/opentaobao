package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest
分销状态查询接口queryProjectStatusByProjectId API请求
alibaba.damai.maitix.opengateway.project.status.query

queryProjectStatusByProjectId */
type AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest struct {
	model.Params
	// 入参dto
	_disProjectStatusQueryParam *DisProjectStatusQueryDto
}

// New
