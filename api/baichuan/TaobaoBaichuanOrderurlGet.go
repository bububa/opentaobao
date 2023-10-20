package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOrderurlGet 百川订单详情
// taobao.baichuan.orderurl.get
//
// 百川订单详情
func TaobaoBaichuanOrderurlGet(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOrderurlGetAPIRequest, session string) (*baichuan.TaobaoBaichuanOrderurlGetAPIResponse, error) {
	var resp baichuan.TaobaoBaichuanOrderurlGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
