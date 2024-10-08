package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerCheckppw 校验支付密码
// alibaba.alsc.crm.customer.checkppw
//
// 校验支付密码
func AlibabaAlscCrmCustomerCheckppw(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerCheckppwAPIRequest, resp *alsc.AlibabaAlscCrmCustomerCheckppwAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
