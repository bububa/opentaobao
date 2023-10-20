package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemsizemappingtemplateslist 获取天猫商品尺码表模板列表
// tmall.item.sizemapping.templates.list
//
// 获取所有尺码表模板列表。
func Tmallitemsizemappingtemplateslist(clt *core.SDKClient, req *product.TmallitemsizemappingtemplateslistAPIRequest, session string) (*product.TmallitemsizemappingtemplateslistAPIResponse, error) {
	var resp product.TmallitemsizemappingtemplateslistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
