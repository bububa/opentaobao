package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoWlbOrderJzConsign 家装发货接口
// taobao.wlb.order.jz.consign
//
// 家装类订单使用该接口发货
func TaobaoWlbOrderJzConsign(clt *core.SDKClient, req *tblogistics.TaobaoWlbOrderJzConsignAPIRequest, session string) (*tblogistics.TaobaoWlbOrderJzConsignAPIResponse, error) {
	var resp tblogistics.TaobaoWlbOrderJzConsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
