package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TmallItemSchemaUpdate
天猫根据规则编辑商品
tmall.item.schema.update

天猫根据规则编辑商品 */
func TmallItemSchemaUpdate(clt *core.SDKClient, req *product.TmallItemSchemaUpdateAPIRequest, session string) (*product.TmallItemSchemaUpdateAPIResponse, error) {
	var resp product.TmallItemSchemaUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
