package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCrowdCrowdFindcrowdgroups 获取人群组
// taobao.onebp.dkx.crowd.crowd.findcrowdgroups
//
// 获取人群组
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{ &#34;market_scene&#34;: &#34;ad_strategy_laxin&#34;}
func TaobaoOnebpDkxCrowdCrowdFindcrowdgroups(clt *core.SDKClient, req *scs.TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest, resp *scs.TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
