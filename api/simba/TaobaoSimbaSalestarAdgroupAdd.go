package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarAdgroupAdd (新)创建一个推广组
// taobao.simba.salestar.adgroup.add
//
// 创建一个推广组
func TaobaoSimbaSalestarAdgroupAdd(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarAdgroupAddAPIRequest, resp *simba.TaobaoSimbaSalestarAdgroupAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
