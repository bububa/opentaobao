package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoWlbOrderJzQuery 家装业务查询物流公司api
// taobao.wlb.order.jz.query
//
// 家装业务查询物流公司api
func TaobaoWlbOrderJzQuery(clt *core.SDKClient, req *tblogistics.TaobaoWlbOrderJzQueryAPIRequest, session string) (*tblogistics.TaobaoWlbOrderJzQueryAPIResponse, error) {
	var resp tblogistics.TaobaoWlbOrderJzQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
