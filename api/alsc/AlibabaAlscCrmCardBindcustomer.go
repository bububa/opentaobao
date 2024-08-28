package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardBindcustomer 卡号绑定顾客
// alibaba.alsc.crm.card.bindcustomer
//
// 为卡号绑定顾客
func AlibabaAlscCrmCardBindcustomer(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardBindcustomerAPIRequest, resp *alsc.AlibabaAlscCrmCardBindcustomerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
