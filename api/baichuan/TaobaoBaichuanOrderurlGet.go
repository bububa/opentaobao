package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOrderurlGet 百川订单详情
// taobao.baichuan.orderurl.get
//
// 百川订单详情
func TaobaoBaichuanOrderurlGet(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanOrderurlGetAPIRequest, resp *baichuan.TaobaoBaichuanOrderurlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
