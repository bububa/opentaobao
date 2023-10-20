package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaobanamadpcitemupdate 编辑商品
// taobao.banamadpc.item.update
//
// 巴拿马供应商通过此接口编辑商品
func Taobaobanamadpcitemupdate(clt *core.SDKClient, req *product.TaobaobanamadpcitemupdateAPIRequest, session string) (*product.TaobaobanamadpcitemupdateAPIResponse, error) {
	var resp product.TaobaobanamadpcitemupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
