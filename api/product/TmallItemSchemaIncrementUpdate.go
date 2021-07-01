package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TmallItemSchemaIncrementUpdate
天猫根据规则增量更新商品
tmall.item.schema.increment.update

增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名） */
func TmallItemSchemaIncrementUpdate(clt *core.SDKClient, req *product.TmallItemSchemaIncrementUpdateAPIRequest, session string) (*product.TmallItemSchemaIncrementUpdateAPIResponse, error) {
	var resp product.TmallItemSchemaIncrementUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
