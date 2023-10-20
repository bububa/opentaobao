package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoBanamadpcItemAdd 新发商品
// taobao.banamadpc.item.add
//
// 巴拿马供应商通过此接口新发商品
func TaobaoBanamadpcItemAdd(clt *core.SDKClient, req *product.TaobaoBanamadpcItemAddAPIRequest, resp *product.TaobaoBanamadpcItemAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
