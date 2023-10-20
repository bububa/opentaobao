package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSizemappingTemplateUpdate 更新天猫商品尺码表模板
// tmall.item.sizemapping.template.update
//
// 更新天猫商品尺码表模板
func TmallItemSizemappingTemplateUpdate(clt *core.SDKClient, req *product.TmallItemSizemappingTemplateUpdateAPIRequest, session string) (*product.TmallItemSizemappingTemplateUpdateAPIResponse, error) {
	var resp product.TmallItemSizemappingTemplateUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
