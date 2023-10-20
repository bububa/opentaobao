package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallproductaddschemaget 产品发布规则获取接口
// tmall.product.add.schema.get
//
// 获取用户发布产品的规则
func Tmallproductaddschemaget(clt *core.SDKClient, req *tbitem.TmallproductaddschemagetAPIRequest, session string) (*tbitem.TmallproductaddschemagetAPIResponse, error) {
	var resp tbitem.TmallproductaddschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
