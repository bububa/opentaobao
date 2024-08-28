package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRtrptCreativeGet 获取创意实时报表数据
// taobao.simba.rtrpt.creative.get
//
// 获取创意实时报表数据
func TaobaoSimbaRtrptCreativeGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRtrptCreativeGetAPIRequest, resp *simba.TaobaoSimbaRtrptCreativeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
