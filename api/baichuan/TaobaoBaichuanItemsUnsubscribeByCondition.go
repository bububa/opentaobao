package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanitemsunsubscribebycondition 根据条件删除订阅关系
// taobao.baichuan.items.unsubscribe.by.condition
//
// 根据条件删除订阅关系
func Taobaobaichuanitemsunsubscribebycondition(clt *core.SDKClient, req *baichuan.TaobaobaichuanitemsunsubscribebyconditionAPIRequest, session string) (*baichuan.TaobaobaichuanitemsunsubscribebyconditionAPIResponse, error) {
	var resp baichuan.TaobaobaichuanitemsunsubscribebyconditionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
