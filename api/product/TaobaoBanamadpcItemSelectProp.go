package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaobanamadpcitemselectprop 获取子属性
// taobao.banamadpc.item.select.prop
//
// 巴拿马供应商通过此接口获取子属性
func Taobaobanamadpcitemselectprop(clt *core.SDKClient, req *product.TaobaobanamadpcitemselectpropAPIRequest, session string) (*product.TaobaobanamadpcitemselectpropAPIResponse, error) {
	var resp product.TaobaobanamadpcitemselectpropAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
