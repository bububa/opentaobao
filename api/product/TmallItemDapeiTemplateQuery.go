package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemDapeiTemplateQuery 搭配查询接口
// tmall.item.dapei.template.query
//
// 根据条件获取搭配内容
func TmallItemDapeiTemplateQuery(clt *core.SDKClient, req *product.TmallItemDapeiTemplateQueryAPIRequest, session string) (*product.TmallItemDapeiTemplateQueryAPIResponse, error) {
	var resp product.TmallItemDapeiTemplateQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
