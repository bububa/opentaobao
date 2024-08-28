package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardRefuse 买家拒收
// tmall.servicecenter.workcard.refuse
//
// 买家拒收通知接口
func TmallServicecenterWorkcardRefuse(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardRefuseAPIRequest, resp *tmallsc.TmallServicecenterWorkcardRefuseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
