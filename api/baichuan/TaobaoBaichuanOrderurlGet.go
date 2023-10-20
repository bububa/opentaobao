package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanorderurlget 百川订单详情
// taobao.baichuan.orderurl.get
//
// 百川订单详情
func Taobaobaichuanorderurlget(clt *core.SDKClient, req *baichuan.TaobaobaichuanorderurlgetAPIRequest, session string) (*baichuan.TaobaobaichuanorderurlgetAPIResponse, error) {
	var resp baichuan.TaobaobaichuanorderurlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
