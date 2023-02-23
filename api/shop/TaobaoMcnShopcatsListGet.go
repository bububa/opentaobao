package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoMcnShopcatsListGet 店铺类目清单
// taobao.mcn.shopcats.list.get
//
// 无需授权; 获取前台展示的店铺类目;
func TaobaoMcnShopcatsListGet(clt *core.SDKClient, req *shop.TaobaoMcnShopcatsListGetAPIRequest, session string) (*shop.TaobaoMcnShopcatsListGetAPIResponse, error) {
	var resp shop.TaobaoMcnShopcatsListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
