package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TmallItemSchemaAdd
天猫根据规则发布商品
tmall.item.schema.add

天猫TopSchema发布商品。 */
func TmallItemSchemaAdd(clt *core.SDKClient, req *product.TmallItemSchemaAddAPIRequest, session string) (*product.TmallItemSchemaAddAPIResponse, error) {
	var resp product.TmallItemSchemaAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
