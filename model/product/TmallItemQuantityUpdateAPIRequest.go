package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemQuantityUpdateAPIRequest
天猫商品/SKU库存更新接口 API请求
tmall.item.quantity.update

天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。 */
type TmallItemQuantityUpdateAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
	_skuQuantities []UpdateSkuQuantity
	// 商品库存更新时候的可选参数
	_options *UpdateItemQuantityOption
	// 商品库存数；增量编辑方式支持正数、负数。（无SKU商品使用这个字段）
	_itemQuantity int64
}

// New
