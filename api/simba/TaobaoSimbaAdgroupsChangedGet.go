package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupsChangedGet 分页获取修改的推广组ID和修改时间
// taobao.simba.adgroups.changed.get
//
// 分页获取修改的推广组ID和修改时间
func TaobaoSimbaAdgroupsChangedGet(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupsChangedGetAPIRequest, resp *simba.TaobaoSimbaAdgroupsChangedGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
