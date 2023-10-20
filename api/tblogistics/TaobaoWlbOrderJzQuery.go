package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoWlbOrderJzQuery 家装业务查询物流公司api
// taobao.wlb.order.jz.query
//
// 家装业务查询物流公司api
func TaobaoWlbOrderJzQuery(clt *core.SDKClient, req *tblogistics.TaobaoWlbOrderJzQueryAPIRequest, resp *tblogistics.TaobaoWlbOrderJzQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
