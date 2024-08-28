package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmVoucherStatusChange 优惠券状态更改
// alibaba.alsc.crm.voucher.status.change
//
// 核销优惠券
func AlibabaAlscCrmVoucherStatusChange(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmVoucherStatusChangeAPIRequest, resp *alsc.AlibabaAlscCrmVoucherStatusChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
