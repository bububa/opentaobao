package alscmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// AlibabaAlscMerchantExtTicketcodeSend 异步发码
// alibaba.alsc.merchant.ext.ticketcode.send
//
// 外部券异步发码
func AlibabaAlscMerchantExtTicketcodeSend(ctx context.Context, clt *core.SDKClient, req *alscmerchant.AlibabaAlscMerchantExtTicketcodeSendAPIRequest, resp *alscmerchant.AlibabaAlscMerchantExtTicketcodeSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
