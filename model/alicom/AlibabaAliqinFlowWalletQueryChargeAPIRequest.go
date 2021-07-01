package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowWalletQueryChargeAPIRequest
查询流量充值状态 API请求
alibaba.aliqin.flow.wallet.query.charge

查询流量充值状态 */
type AlibabaAliqinFlowWalletQueryChargeAPIRequest struct {
	model.Params
	// 唯一流水号
	_outRechargeId string
	// 渠道id
	_channelId string
}

// New
