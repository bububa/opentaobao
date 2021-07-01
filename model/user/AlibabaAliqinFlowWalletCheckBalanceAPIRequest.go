package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowWalletCheckBalanceAPIRequest
商家预存余额检查 API请求
alibaba.aliqin.flow.wallet.check.balance

检查商家CRM预存余额是否足够进行活动 */
type AlibabaAliqinFlowWalletCheckBalanceAPIRequest struct {
	model.Params
	// 检查金额档位id
	_gradeId string
}

// New
