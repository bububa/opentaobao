package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemsSubscribe 百川批量商品订阅
// taobao.baichuan.items.subscribe
//
// 百川批量添加订阅的商品
func TaobaoBaichuanItemsSubscribe(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemsSubscribeAPIRequest, resp *baichuan.TaobaoBaichuanItemsSubscribeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
