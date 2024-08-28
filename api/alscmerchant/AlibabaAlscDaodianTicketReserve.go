package alscmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// AlibabaAlscDaodianTicketReserve 外部券冲正
// alibaba.alsc.daodian.ticket.reserve
//
// 外部券冲正
func AlibabaAlscDaodianTicketReserve(ctx context.Context, clt *core.SDKClient, req *alscmerchant.AlibabaAlscDaodianTicketReserveAPIRequest, resp *alscmerchant.AlibabaAlscDaodianTicketReserveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
