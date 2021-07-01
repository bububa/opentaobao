package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductSkuUpdateAPIRequest
产品sku编辑接口 API请求
taobao.fenxiao.product.sku.update

产品SKU信息更新 */
type TaobaoFenxiaoProductSkuUpdateAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 产品SKU库存
	_quantity int64
	// 采购基准价
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
