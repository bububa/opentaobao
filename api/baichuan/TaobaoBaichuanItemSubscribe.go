package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemSubscribe 单个商品订阅
// taobao.baichuan.item.subscribe
//
// 百川单个商品订阅
func TaobaoBaichuanItemSubscribe(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemSubscribeAPIRequest, resp *baichuan.TaobaoBaichuanItemSubscribeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
