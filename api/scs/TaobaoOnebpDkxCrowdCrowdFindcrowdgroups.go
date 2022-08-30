package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCrowdCrowdFindcrowdgroups 获取人群组
// taobao.onebp.dkx.crowd.crowd.findcrowdgroups
//
// 获取人群组
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{ "market_scene": "ad_strategy_laxin"}
func TaobaoOnebpDkxCrowdCrowdFindcrowdgroups(clt *core.SDKClient, req *scs.TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIRequest, session string) (*scs.TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
