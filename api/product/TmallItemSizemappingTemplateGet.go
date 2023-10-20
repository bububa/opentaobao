package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemsizemappingtemplateget 获取天猫商品尺码表模板
// tmall.item.sizemapping.template.get
//
// 获取天猫商品尺码表模板
func Tmallitemsizemappingtemplateget(clt *core.SDKClient, req *product.TmallitemsizemappingtemplategetAPIRequest, session string) (*product.TmallitemsizemappingtemplategetAPIResponse, error) {
	var resp product.TmallitemsizemappingtemplategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
