package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductSkuAddAPIRequest
产品sku添加接口 API请求
taobao.fenxiao.product.sku.add

添加产品SKU信息 */
type TaobaoFenxiaoProductSkuAddAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// sku产品库存，在0到1000000之间，如果不传，则库存为0
	_quantity int64
	// 采购基准价，最大值1000000000
	_standardPrice string
	// 代销采购价
	_agentCostPrice string
	// sku属性
	_properties string
	// 商家编码
	_skuNumber string
	// 经销采购价
	_dealerCostPrice string
}

// New
