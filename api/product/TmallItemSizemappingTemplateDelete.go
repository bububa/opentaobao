package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemsizemappingtemplatedelete 删除天猫商品尺码表模板
// tmall.item.sizemapping.template.delete
//
// 删除天猫商品尺码表模板
func Tmallitemsizemappingtemplatedelete(clt *core.SDKClient, req *product.TmallitemsizemappingtemplatedeleteAPIRequest, session string) (*product.TmallitemsizemappingtemplatedeleteAPIResponse, error) {
	var resp product.TmallitemsizemappingtemplatedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
