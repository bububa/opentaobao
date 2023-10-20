package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallProductSchemaUpdate 产品更新接口
// tmall.product.schema.update
//
// 产品更新接口
func TmallProductSchemaUpdate(clt *core.SDKClient, req *tbitem.TmallProductSchemaUpdateAPIRequest, session string) (*tbitem.TmallProductSchemaUpdateAPIResponse, error) {
	var resp tbitem.TmallProductSchemaUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
