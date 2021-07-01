package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TmallProductSchemaUpdate
产品更新接口
tmall.product.schema.update

产品更新接口 */
func TmallProductSchemaUpdate(clt *core.SDKClient, req *product.TmallProductSchemaUpdateAPIRequest, session string) (*product.TmallProductSchemaUpdateAPIResponse, error) {
	var resp product.TmallProductSchemaUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
