package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanitemsubscribe 单个商品订阅
// taobao.baichuan.item.subscribe
//
// 百川单个商品订阅
func Taobaobaichuanitemsubscribe(clt *core.SDKClient, req *baichuan.TaobaobaichuanitemsubscribeAPIRequest, session string) (*baichuan.TaobaobaichuanitemsubscribeAPIResponse, error) {
	var resp baichuan.TaobaobaichuanitemsubscribeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
