package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemSchemaIncrementUpdateAPIRequest
天猫根据规则增量更新商品 API请求
tmall.item.schema.increment.update

增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名） */
type TmallItemSchemaIncrementUpdateAPIRequest struct {
	model.Params
	// 需要编辑的商品ID
	_itemId int64
	// 根据tmall.item.increment.update.schema.get生成的商品增量编辑规则入参数据。需要更新的字段，一定要在入参的XML重点update_fields字段中明确指明
	_xmlData string
}

// New
