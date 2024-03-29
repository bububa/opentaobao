package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSizemappingTemplateDelete 删除天猫商品尺码表模板
// tmall.item.sizemapping.template.delete
//
// 删除天猫商品尺码表模板
func TmallItemSizemappingTemplateDelete(clt *core.SDKClient, req *product.TmallItemSizemappingTemplateDeleteAPIRequest, resp *product.TmallItemSizemappingTemplateDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
