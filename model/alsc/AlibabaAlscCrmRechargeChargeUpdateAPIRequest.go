package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeChargeUpdateAPIRequest
储值充值 API请求
alibaba.alsc.crm.recharge.charge.update

顾客储值账户充值 */
type AlibabaAlscCrmRechargeChargeUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramRechargeOpenReq *RechargeOpenReq
}

// New
