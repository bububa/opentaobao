package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemSchemaIncrementUpdate 天猫根据规则增量更新商品
// tmall.item.schema.increment.update
//
// 增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名）
func TmallItemSchemaIncrementUpdate(clt *core.SDKClient, req *tbitem.TmallItemSchemaIncrementUpdateAPIRequest, resp *tbitem.TmallItemSchemaIncrementUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
