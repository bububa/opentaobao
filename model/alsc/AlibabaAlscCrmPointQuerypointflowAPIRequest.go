package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPointQuerypointflowAPIRequest
分页查询积分流水 API请求
alibaba.alsc.crm.point.querypointflow

分页查询积分流水 */
type AlibabaAlscCrmPointQuerypointflowAPIRequest struct {
	model.Params
	// 入参
	_paramPageQueryPointFlowOpenReq *PageQueryPointFlowOpenReq
}

// New
