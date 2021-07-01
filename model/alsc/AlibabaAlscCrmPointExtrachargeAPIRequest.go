package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPointExtrachargeAPIRequest
积分补录 API请求
alibaba.alsc.crm.point.extracharge

积分补录 */
type AlibabaAlscCrmPointExtrachargeAPIRequest struct {
	model.Params
	// 入参
	_paramExtraChargePointOpenReq *ExtraChargePointOpenReq
}

// New
