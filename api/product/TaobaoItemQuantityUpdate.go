package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemQuantityUpdate 宝贝/SKU库存修改
// taobao.item.quantity.update
//
// 提供按照全量或增量形式修改宝贝/SKU库存的功能
func TaobaoItemQuantityUpdate(clt *core.SDKClient, req *product.TaobaoItemQuantityUpdateAPIRequest, session string) (*product.TaobaoItemQuantityUpdateAPIResponse, error) {
	var resp product.TaobaoItemQuantityUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
