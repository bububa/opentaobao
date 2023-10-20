package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemsUnsubscribe 批量删除商品订阅
// taobao.baichuan.items.unsubscribe
//
// 批量删除商品订阅
func TaobaoBaichuanItemsUnsubscribe(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemsUnsubscribeAPIRequest, resp *baichuan.TaobaoBaichuanItemsUnsubscribeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
