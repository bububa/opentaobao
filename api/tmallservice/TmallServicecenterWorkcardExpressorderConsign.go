package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardExpressorderConsign 物流订单呼叫运力
// tmall.servicecenter.workcard.expressorder.consign
//
// 天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力
func TmallServicecenterWorkcardExpressorderConsign(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardExpressorderConsignAPIRequest, resp *tmallservice.TmallServicecenterWorkcardExpressorderConsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
