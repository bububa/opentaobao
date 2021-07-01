package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosflowWorkQueryvariablesAPIRequest
获取指定流程上下文参数 API请求
alibaba.mosflow.work.queryvariables

业务查询指定流程上下文内容 */
type AlibabaMosflowWorkQueryvariablesAPIRequest struct {
	model.Params
	// 流程实例ID
	_processInstanceId string
}

// New
