package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCrowdCrowdCoverage 获取人数预估
// taobao.onebp.dkx.crowd.crowd.coverage
//
// 获取人数预估，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func TaobaoOnebpDkxCrowdCrowdCoverage(clt *core.SDKClient, req *scs.TaobaoOnebpDkxCrowdCrowdCoverageAPIRequest, resp *scs.TaobaoOnebpDkxCrowdCrowdCoverageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
