package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowWalletChargeAPIRequest
流量直充 API请求
alibaba.aliqin.flow.wallet.charge

流量直充 */
type AlibabaAliqinFlowWalletChargeAPIRequest struct {
	model.Params
	// 充值号码
	_phoneNum string
	// 原因
	_reason string
	// 档位id
	_gradeId string
	// 唯一流水号
	_outRechargeId string
	// 渠道id
	_channelId string
}

// New
