package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoMiniappItemsGet 批量获取商品信息
// taobao.miniapp.items.get
//
// 获取商品公开属性，只允许在商家应用环境中使用
func TaobaoMiniappItemsGet(clt *core.SDKClient, req *product.TaobaoMiniappItemsGetAPIRequest, session string) (*product.TaobaoMiniappItemsGetAPIResponse, error) {
	var resp product.TaobaoMiniappItemsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
