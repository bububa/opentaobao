package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPointConsumepointAPIRequest
积分抵现 API请求
alibaba.alsc.crm.point.consumepoint

积分抵现 */
type AlibabaAlscCrmPointConsumepointAPIRequest struct {
	model.Params
	// 入参
	_paramConsumePointOpenReq *ConsumePointOpenReq
}

// New
