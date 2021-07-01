package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemUpdateSchemaGetAPIRequest
天猫编辑商品规则获取 API请求
tmall.item.update.schema.get

Schema方式编辑天猫商品时，编辑商品规则获取 */
type TmallItemUpdateSchemaGetAPIRequest struct {
	model.Params
	// 需要编辑的商品ID
	_itemId int64
	// 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。
	_categoryId int64
	// 商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。
	_productId int64
}

// New
