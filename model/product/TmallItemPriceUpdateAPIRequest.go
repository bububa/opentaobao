package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemPriceUpdateAPIRequest
天猫商品/SKU价格更新接口 API请求
tmall.item.price.update

天猫商品/SKU价格更新接口，支持商品、SKU价格同时更新，支持同一商品下的SKU批量更新。 */
type TmallItemPriceUpdateAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 被更新商品价格
	_itemPrice float64
	// 更新SKU价格时候的SKU价格对象；如果没有SKU或者不更新SKU价格，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
	_skuPrices []UpdateSkuPrice
	// 商品价格更新时候的可选参数
	_options *UpdateItemPriceOption
}

// New
