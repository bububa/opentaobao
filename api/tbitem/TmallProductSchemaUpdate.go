package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallproductschemaupdate 产品更新接口
// tmall.product.schema.update
//
// 产品更新接口
func Tmallproductschemaupdate(clt *core.SDKClient, req *tbitem.TmallproductschemaupdateAPIRequest, session string) (*tbitem.TmallproductschemaupdateAPIResponse, error) {
	var resp tbitem.TmallproductschemaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
