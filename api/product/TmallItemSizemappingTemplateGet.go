package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSizemappingTemplateGet 获取天猫商品尺码表模板
// tmall.item.sizemapping.template.get
//
// 获取天猫商品尺码表模板
func TmallItemSizemappingTemplateGet(clt *core.SDKClient, req *product.TmallItemSizemappingTemplateGetAPIRequest, resp *product.TmallItemSizemappingTemplateGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
