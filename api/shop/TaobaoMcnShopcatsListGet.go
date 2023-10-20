package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Taobaomcnshopcatslistget 店铺类目清单
// taobao.mcn.shopcats.list.get
//
// 无需授权; 获取前台展示的店铺类目;
func Taobaomcnshopcatslistget(clt *core.SDKClient, req *shop.TaobaomcnshopcatslistgetAPIRequest, session string) (*shop.TaobaomcnshopcatslistgetAPIResponse, error) {
	var resp shop.TaobaomcnshopcatslistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
