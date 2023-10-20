package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTopEventSubscriptionQuery 三方事件订阅查询
// taobao.top.event.subscription.query
//
// 三方事件订阅查询
func TaobaoTopEventSubscriptionQuery(clt *core.SDKClient, req *util.TaobaoTopEventSubscriptionQueryAPIRequest, session string) (*util.TaobaoTopEventSubscriptionQueryAPIResponse, error) {
	var resp util.TaobaoTopEventSubscriptionQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
