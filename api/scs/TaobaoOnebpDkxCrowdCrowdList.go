package scs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCrowdCrowdList 获取人群信息列表
// taobao.onebp.dkx.crowd.crowd.list
//
// 获取人群信息列表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func TaobaoOnebpDkxCrowdCrowdList(ctx context.Context, clt *core.SDKClient, req *scs.TaobaoOnebpDkxCrowdCrowdListAPIRequest, resp *scs.TaobaoOnebpDkxCrowdCrowdListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
