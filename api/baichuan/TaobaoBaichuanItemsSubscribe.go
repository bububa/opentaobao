package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanitemssubscribe 百川批量商品订阅
// taobao.baichuan.items.subscribe
//
// 百川批量添加订阅的商品
func Taobaobaichuanitemssubscribe(clt *core.SDKClient, req *baichuan.TaobaobaichuanitemssubscribeAPIRequest, session string) (*baichuan.TaobaobaichuanitemssubscribeAPIResponse, error) {
	var resp baichuan.TaobaobaichuanitemssubscribeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
