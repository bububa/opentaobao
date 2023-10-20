package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallproductschemaadd 使用Schema文件发布一个产品
// tmall.product.schema.add
//
// Schema体系发布一个产品
func Tmallproductschemaadd(clt *core.SDKClient, req *tbitem.TmallproductschemaaddAPIRequest, session string) (*tbitem.TmallproductschemaaddAPIResponse, error) {
	var resp tbitem.TmallproductschemaaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
