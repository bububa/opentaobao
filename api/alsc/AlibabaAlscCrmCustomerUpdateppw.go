package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerUpdateppw 修改支付密码
// alibaba.alsc.crm.customer.updateppw
//
// 修改支付密码
func AlibabaAlscCrmCustomerUpdateppw(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerUpdateppwAPIRequest, resp *alsc.AlibabaAlscCrmCustomerUpdateppwAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
