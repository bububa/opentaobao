package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest
储值账户充值前校验 API请求
alibaba.alsc.crm.recharge.chargeprecheck.get

储值账户充值前校验接口 */
type AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest struct {
	model.Params
	// 入参
	_paramChargePreCheckOpenReq *ChargePreCheckOpenReq
}

// New
