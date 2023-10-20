package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupidsDeletedGet 获取删除的推广组ID
// taobao.simba.adgroupids.deleted.get
//
// 获取删除的推广组ID
func TaobaoSimbaAdgroupidsDeletedGet(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupidsDeletedGetAPIRequest, resp *simba.TaobaoSimbaAdgroupidsDeletedGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
