package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Taobaoshopcatslistget 获取前台展示的店铺类目
// taobao.shopcats.list.get
//
// 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
func Taobaoshopcatslistget(clt *core.SDKClient, req *shop.TaobaoshopcatslistgetAPIRequest, session string) (*shop.TaobaoshopcatslistgetAPIResponse, error) {
	var resp shop.TaobaoshopcatslistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
