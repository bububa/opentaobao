package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallproductmatchschemaget 获取匹配产品规则
// tmall.product.match.schema.get
//
// ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
func Tmallproductmatchschemaget(clt *core.SDKClient, req *tbitem.TmallproductmatchschemagetAPIRequest, session string) (*tbitem.TmallproductmatchschemagetAPIResponse, error) {
	var resp tbitem.TmallproductmatchschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
