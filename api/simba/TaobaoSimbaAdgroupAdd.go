package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupAdd 创建一个推广组
// taobao.simba.adgroup.add
//
// 创建一个推广组
func TaobaoSimbaAdgroupAdd(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupAddAPIRequest, resp *simba.TaobaoSimbaAdgroupAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
