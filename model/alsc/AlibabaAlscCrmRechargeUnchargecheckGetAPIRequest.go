package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest
储值账户退充值校验 API请求
alibaba.alsc.crm.recharge.unchargecheck.get

储值账户退充值校验接口 */
type AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest struct {
	model.Params
	// 入参
	_paramUnchargeCheckOpenReq *UnchargeCheckOpenReq
}

// New
