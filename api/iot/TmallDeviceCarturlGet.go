package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Tmalldevicecarturlget 添加商品到购物车
// tmall.device.carturl.get
//
// 获取二维码，支持添加商品到购物车
func Tmalldevicecarturlget(clt *core.SDKClient, req *iot.TmalldevicecarturlgetAPIRequest, session string) (*iot.TmalldevicecarturlgetAPIResponse, error) {
	var resp iot.TmalldevicecarturlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
