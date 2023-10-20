package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemSubscribe 单个商品订阅
// taobao.baichuan.item.subscribe
//
// 百川单个商品订阅
func TaobaoBaichuanItemSubscribe(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemSubscribeAPIRequest, resp *baichuan.TaobaoBaichuanItemSubscribeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
