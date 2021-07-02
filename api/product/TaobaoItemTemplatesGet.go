package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemTemplatesGet 获取用户宝贝详情页模板名称
// taobao.item.templates.get
//
// 查询当前登录用户的店铺的宝贝详情页的模板名称
func TaobaoItemTemplatesGet(clt *core.SDKClient, req *product.TaobaoItemTemplatesGetAPIRequest, session string) (*product.TaobaoItemTemplatesGetAPIResponse, error) {
	var resp product.TaobaoItemTemplatesGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
