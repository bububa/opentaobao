package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpmaterialshopget 获取店铺信息
// taobao.universalbp.material.shop.get
//
// 获取店铺信息
func Taobaouniversalbpmaterialshopget(clt *core.SDKClient, req *simba.TaobaouniversalbpmaterialshopgetAPIRequest, session string) (*simba.TaobaouniversalbpmaterialshopgetAPIResponse, error) {
	var resp simba.TaobaouniversalbpmaterialshopgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
