package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemSchemaUpdateAPIRequest
天猫根据规则编辑商品 API请求
tmall.item.schema.update

天猫根据规则编辑商品 */
type TmallItemSchemaUpdateAPIRequest struct {
	model.Params
	// 需要编辑的商品ID
	_itemId int64
	// 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求不需要填写
	_categoryId int64
	// 商品发布的目标product_id。如果没有切换类目或者切换产品的需求，参数不用填写
	_productId int64
	// 根据tmall.item.update.schema.get生成的商品编辑规则入参数据
	_xmlData string
}

// New
