package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradeorderskuupdateAPIRequest 更新交易的销售属性 API请求
// taobao.trade.ordersku.update
//
// 只能更新发货前子订单的销售属性 &lt;br/&gt;只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 &lt;br/&gt;必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先
type TaobaotradeorderskuupdateAPIRequest struct {
	model.Params
	// 销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
	_skuProps string
	// 子订单编号（对于单笔订单的交易可以传交易编号）。
	_oid int64
	// 销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
	_skuId int64
}

// NewTaobaotradeorderskuupdateRequest 初始化TaobaotradeorderskuupdateAPIRequest对象
func NewTaobaotradeorderskuupdateRequest() *TaobaotradeorderskuupdateAPIRequest {
	return &TaobaotradeorderskuupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradeorderskuupdateAPIRequest) GetApiMethodName() string {
	return "taobao.trade.ordersku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradeorderskuupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradeorderskuupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuProps is SkuProps Setter
// 销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
func (r *TaobaotradeorderskuupdateAPIRequest) SetSkuProps(_skuProps string) error {
	r._skuProps = _skuProps
	r.Set("sku_props", _skuProps)
	return nil
}

// GetSkuProps SkuProps Getter
func (r TaobaotradeorderskuupdateAPIRequest) GetSkuProps() string {
	return r._skuProps
}

// SetOid is Oid Setter
// 子订单编号（对于单笔订单的交易可以传交易编号）。
func (r *TaobaotradeorderskuupdateAPIRequest) SetOid(_oid int64) error {
	r._oid = _oid
	r.Set("oid", _oid)
	return nil
}

// GetOid Oid Getter
func (r TaobaotradeorderskuupdateAPIRequest) GetOid() int64 {
	return r._oid
}

// SetSkuId is SkuId Setter
// 销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
func (r *TaobaotradeorderskuupdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaotradeorderskuupdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}
