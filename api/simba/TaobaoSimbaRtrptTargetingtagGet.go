package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRtrptTargetingtagGet 搜索人群实时报表
// taobao.simba.rtrpt.targetingtag.get
//
// 获取搜搜人群实时报表
func TaobaoSimbaRtrptTargetingtagGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRtrptTargetingtagGetAPIRequest, resp *simba.TaobaoSimbaRtrptTargetingtagGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
