package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

/* TaobaoWlbOrderJzpartnerQuery
查询家装服务商列表
taobao.wlb.order.jzpartner.query

为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口 */
func TaobaoWlbOrderJzpartnerQuery(clt *core.SDKClient, req *wlb.TaobaoWlbOrderJzpartnerQueryAPIRequest, session string) (*wlb.TaobaoWlbOrderJzpartnerQueryAPIResponse, error) {
	var resp wlb.TaobaoWlbOrderJzpartnerQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
