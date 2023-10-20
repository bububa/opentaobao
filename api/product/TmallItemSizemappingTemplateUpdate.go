package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemsizemappingtemplateupdate 更新天猫商品尺码表模板
// tmall.item.sizemapping.template.update
//
// 更新天猫商品尺码表模板
func Tmallitemsizemappingtemplateupdate(clt *core.SDKClient, req *product.TmallitemsizemappingtemplateupdateAPIRequest, session string) (*product.TmallitemsizemappingtemplateupdateAPIResponse, error) {
	var resp product.TmallitemsizemappingtemplateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
