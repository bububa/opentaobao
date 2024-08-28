package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerGet 查询顾客详情
// alibaba.alsc.crm.customer.get
//
// 查询顾客详情
func AlibabaAlscCrmCustomerGet(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerGetAPIRequest, resp *alsc.AlibabaAlscCrmCustomerGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
