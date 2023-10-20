package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbSubscriptionQuery 查询商家定购的所有服务
// taobao.wlb.subscription.query
//
// 查询商家定购的所有服务,可通过入参状态来筛选
func TaobaoWlbSubscriptionQuery(clt *core.SDKClient, req *wlb.TaobaoWlbSubscriptionQueryAPIRequest, session string) (*wlb.TaobaoWlbSubscriptionQueryAPIResponse, error) {
	var resp wlb.TaobaoWlbSubscriptionQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
