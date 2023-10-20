package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOrderurlGet 百川订单详情
// taobao.baichuan.orderurl.get
//
// 百川订单详情
func TaobaoBaichuanOrderurlGet(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOrderurlGetAPIRequest, resp *baichuan.TaobaoBaichuanOrderurlGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
