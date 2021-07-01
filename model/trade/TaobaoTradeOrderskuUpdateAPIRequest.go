package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeOrderskuUpdateAPIRequest
更新交易的销售属性 API请求
taobao.trade.ordersku.update

只能更新发货前子订单的销售属性 <br/>只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 <br/>必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先 */
type TaobaoTradeOrderskuUpdateAPIRequest struct {
	model.Params
	// 子订单编号（对于单笔订单的交易可以传交易编号）。
	_oid int64
	// 销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
	_skuId int64
	// 销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
	_skuProps string
}

// New
