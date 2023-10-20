package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupMobilediscountDelete 批量删除adgroup的移动溢价
// taobao.simba.adgroup.mobilediscount.delete
//
// 批量删除adgroup的移动溢价
func TaobaoSimbaAdgroupMobilediscountDelete(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest, resp *simba.TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
