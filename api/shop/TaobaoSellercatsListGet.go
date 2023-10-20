package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Taobaosellercatslistget 获取前台展示的店铺内卖家自定义商品类目
// taobao.sellercats.list.get
//
// 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
func Taobaosellercatslistget(clt *core.SDKClient, req *shop.TaobaosellercatslistgetAPIRequest, session string) (*shop.TaobaosellercatslistgetAPIResponse, error) {
	var resp shop.TaobaosellercatslistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
