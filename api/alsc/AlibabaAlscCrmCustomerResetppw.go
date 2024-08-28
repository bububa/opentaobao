package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerResetppw 重置支付密码
// alibaba.alsc.crm.customer.resetppw
//
// 重置支付密码
func AlibabaAlscCrmCustomerResetppw(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerResetppwAPIRequest, resp *alsc.AlibabaAlscCrmCustomerResetppwAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
