package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptTargetingtagbaseGet 定向基础报表
// taobao.simba.rpt.targetingtagbase.get
//
// 获取定向基础报表
func TaobaoSimbaRptTargetingtagbaseGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptTargetingtagbaseGetAPIRequest, resp *simba.TaobaoSimbaRptTargetingtagbaseGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
