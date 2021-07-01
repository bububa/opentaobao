package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest
充值退款 API请求
alibaba.alsc.crm.recharge.uncharge.update

充值退款 */
type AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramUnchargeOpenReq *UnchargeOpenReq
}

// New
