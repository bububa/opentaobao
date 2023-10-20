package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallProductSpecAdd 添加产品规格
// tmall.product.spec.add
//
// 增加产品规格
func TmallProductSpecAdd(clt *core.SDKClient, req *product.TmallProductSpecAddAPIRequest, resp *product.TmallProductSpecAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
