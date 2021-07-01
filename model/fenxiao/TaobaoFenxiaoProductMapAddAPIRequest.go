package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductMapAddAPIRequest
创建分销和后端商品映射关系 API请求
taobao.fenxiao.product.map.add

创建分销和供应链商品映射关系。 */
type TaobaoFenxiaoProductMapAddAPIRequest struct {
	model.Params
	// 分销产品id。
	_productId int64
	// 后端商品id（如果当前分销产品没有sku和后端商品时需要指定）。
	_scItemId int64
	// 分销产品的sku id。逗号分隔，顺序需要保证与sc_item_ids一致（没有sku就不传）。
	_skuIds string
	// 在有sku的情况下，与各个sku对应的后端商品id列表。逗号分隔，顺序需要保证与sku_ids一致。
	_scItemIds string
	// 是否需要校验商家编码，true不校验，false校验。
	_notCheckOuterCode bool
}

// New
