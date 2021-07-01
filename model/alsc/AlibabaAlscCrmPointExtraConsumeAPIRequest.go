package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPointExtraConsumeAPIRequest
积分补扣 API请求
alibaba.alsc.crm.point.extra.consume

积分补扣 */
type AlibabaAlscCrmPointExtraConsumeAPIRequest struct {
	model.Params
	// 入参
	_paramExtraConsumePointOpenReq *ExtraConsumePointOpenReq
}

// New
