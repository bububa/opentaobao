package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TmallProductSchemaGet
产品信息获取schema获取
tmall.product.schema.get

产品信息获取接口schema形式返回 */
func TmallProductSchemaGet(clt *core.SDKClient, req *product.TmallProductSchemaGetAPIRequest, session string) (*product.TmallProductSchemaGetAPIResponse, error) {
	var resp product.TmallProductSchemaGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
