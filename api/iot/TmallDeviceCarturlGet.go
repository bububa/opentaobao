package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TmallDeviceCarturlGet 添加商品到购物车
// tmall.device.carturl.get
//
// 获取二维码，支持添加商品到购物车
func TmallDeviceCarturlGet(clt *core.SDKClient, req *iot.TmallDeviceCarturlGetAPIRequest, resp *iot.TmallDeviceCarturlGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
