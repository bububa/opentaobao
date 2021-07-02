package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemSkuAdd 添加SKU
// taobao.item.sku.add
//
// 新增一个sku到num_iid指定的商品中 <br/>传入的iid所对应的商品必须属于当前会话的用户
func TaobaoItemSkuAdd(clt *core.SDKClient, req *product.TaobaoItemSkuAddAPIRequest, session string) (*product.TaobaoItemSkuAddAPIResponse, error) {
	var resp product.TaobaoItemSkuAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
