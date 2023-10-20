package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupDelete 删除一个推广组
// taobao.simba.adgroup.delete
//
// 删除一个推广组
func TaobaoSimbaAdgroupDelete(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupDeleteAPIRequest, resp *simba.TaobaoSimbaAdgroupDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
