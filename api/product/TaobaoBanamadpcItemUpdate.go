package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoBanamadpcItemUpdate 编辑商品
// taobao.banamadpc.item.update
//
// 巴拿马供应商通过此接口编辑商品
func TaobaoBanamadpcItemUpdate(clt *core.SDKClient, req *product.TaobaoBanamadpcItemUpdateAPIRequest, resp *product.TaobaoBanamadpcItemUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
