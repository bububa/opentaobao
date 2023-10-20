package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCrowdFindlist 查询人群绑定列表
// taobao.universalbp.crowd.findlist
//
// 查询计划和单元上绑定的人群列表
func TaobaoUniversalbpCrowdFindlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpCrowdFindlistAPIRequest, resp *simba.TaobaoUniversalbpCrowdFindlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
