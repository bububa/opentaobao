package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallproductschemaget 产品信息获取schema获取
// tmall.product.schema.get
//
// 产品信息获取接口schema形式返回
func Tmallproductschemaget(clt *core.SDKClient, req *tbitem.TmallproductschemagetAPIRequest, session string) (*tbitem.TmallproductschemagetAPIResponse, error) {
	var resp tbitem.TmallproductschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
