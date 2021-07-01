package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeQryruleAPIRequest
储值规则下行 API请求
alibaba.alsc.crm.recharge.qryrule

储值规则下行 */
type AlibabaAlscCrmRechargeQryruleAPIRequest struct {
	model.Params
	// 请求对象
	_paramPullRechargeRuleByShopReq *PullRechargeRuleByShopReq
}

// New
