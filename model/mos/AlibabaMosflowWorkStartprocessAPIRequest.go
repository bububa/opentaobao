package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosflowWorkStartprocessAPIRequest
发起流程 API请求
alibaba.mosflow.work.startprocess

业务发起流程审批 */
type AlibabaMosflowWorkStartprocessAPIRequest struct {
	model.Params
	// 参数二:额外业务参数,Map的JSON串
	_variables string
	// 流程必传参数
	_parameterEntity *ParameterEntity
}

// New
