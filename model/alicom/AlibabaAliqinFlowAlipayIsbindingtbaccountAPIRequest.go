package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest
判断支付宝用户是否绑定淘宝账号 API请求
alibaba.aliqin.flow.alipay.isbindingtbaccount

判断支付宝用户是否绑定淘宝账号 */
type AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest struct {
	model.Params
	// 支付宝ID
	_alipayId string
}

// New
