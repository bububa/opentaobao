package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowWalletSignAPIRequest
流量平台用户签约情况查询 API请求
alibaba.aliqin.flow.wallet.sign

流量平台用户签约情况查询 */
type AlibabaAliqinFlowWalletSignAPIRequest struct {
	model.Params
	// 用户昵称
	_userNick string
}

// New
