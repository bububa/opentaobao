package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowAlipayIsbindingtbaccount 判断支付宝用户是否绑定淘宝账号
// alibaba.aliqin.flow.alipay.isbindingtbaccount
//
// 判断支付宝用户是否绑定淘宝账号
func AlibabaAliqinFlowAlipayIsbindingtbaccount(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest, resp *alicom.AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
