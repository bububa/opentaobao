package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemVipAddSchemaGet vip商家发布商品的获取规则接口
// tmall.item.vip.add.schema.get
//
// 获取vip商家发布商品的规则
func TmallItemVipAddSchemaGet(clt *core.SDKClient, req *product.TmallItemVipAddSchemaGetAPIRequest, resp *product.TmallItemVipAddSchemaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
