package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPointReversepointAPIRequest
积分消费回退 API请求
alibaba.alsc.crm.point.reversepoint

积分消费回退 */
type AlibabaAlscCrmPointReversepointAPIRequest struct {
	model.Params
	// 入参
	_paramReverseConsumePointOpenReq *ReverseConsumePointOpenReq
}

// New
