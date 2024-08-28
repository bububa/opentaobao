package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterReservecondUpdate 主动预约条件更新
// tmall.servicecenter.reservecond.update
//
// 1、设置主动预约开通条件
func TmallServicecenterReservecondUpdate(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterReservecondUpdateAPIRequest, resp *tmallsc.TmallServicecenterReservecondUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
