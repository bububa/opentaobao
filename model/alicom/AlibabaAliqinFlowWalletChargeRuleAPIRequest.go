package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowWalletChargeRuleAPIRequest
流量钱包直充（根据号码归属地省份路由） API请求
alibaba.aliqin.flow.wallet.charge.rule

流量钱包直充（根据号码归属地省份路由） */
type AlibabaAliqinFlowWalletChargeRuleAPIRequest struct {
	model.Params
	// 号码
	_phoneNum string
	// 原因
	_reason string
	// 档位id
	_gradeId string
	// 唯一流水号
	_outRechargeId string
	// 渠道id（运营分配）
	_channelId string
}

// New
