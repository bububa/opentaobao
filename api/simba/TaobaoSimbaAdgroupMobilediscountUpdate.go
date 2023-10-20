package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupMobilediscountUpdate 对推广组进行单独移动溢价
// taobao.simba.adgroup.mobilediscount.update
//
// 对推广组进行单独移动溢价
func TaobaoSimbaAdgroupMobilediscountUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest, resp *simba.TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
