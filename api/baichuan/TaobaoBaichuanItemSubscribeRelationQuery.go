package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanitemsubscriberelationquery 查询单个订阅关系
// taobao.baichuan.item.subscribe.relation.query
//
// 查询单个订阅关系
func Taobaobaichuanitemsubscriberelationquery(clt *core.SDKClient, req *baichuan.TaobaobaichuanitemsubscriberelationqueryAPIRequest, session string) (*baichuan.TaobaobaichuanitemsubscriberelationqueryAPIResponse, error) {
	var resp baichuan.TaobaobaichuanitemsubscriberelationqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
