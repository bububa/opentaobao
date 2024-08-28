package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoFuwuSpConfirmApply 内购服务确认单申请接口
// taobao.fuwu.sp.confirm.apply
//
// isv能通过该接口发起确认申请单
func TaobaoFuwuSpConfirmApply(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoFuwuSpConfirmApplyAPIRequest, resp *servicecenter.TaobaoFuwuSpConfirmApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
