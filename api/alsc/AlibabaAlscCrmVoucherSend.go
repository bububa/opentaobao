package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmVoucherSend 发送券给指定用户
// alibaba.alsc.crm.voucher.send
//
// 发送券给指定用户
func AlibabaAlscCrmVoucherSend(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmVoucherSendAPIRequest, resp *alsc.AlibabaAlscCrmVoucherSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
