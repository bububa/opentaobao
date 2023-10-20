package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanitemsunsubscribe 批量删除商品订阅
// taobao.baichuan.items.unsubscribe
//
// 批量删除商品订阅
func Taobaobaichuanitemsunsubscribe(clt *core.SDKClient, req *baichuan.TaobaobaichuanitemsunsubscribeAPIRequest, session string) (*baichuan.TaobaobaichuanitemsunsubscribeAPIResponse, error) {
	var resp baichuan.TaobaobaichuanitemsunsubscribeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
