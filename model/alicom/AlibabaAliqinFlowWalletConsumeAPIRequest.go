package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowWalletConsumeAPIRequest
流量扣减 API请求
alibaba.aliqin.flow.wallet.consume

流量钱包流量扣减接口 */
type AlibabaAliqinFlowWalletConsumeAPIRequest struct {
	model.Params
	// 扣减流量值
	_flow int64
	// 扣减流水号
	_serialNo string
	// 扣减原因
	_reason string
	// 备注
	_remark string
}

// New
