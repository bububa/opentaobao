package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeOrderskuUpdateAPIRequest 更新交易的销售属性 API请求
// taobao.trade.ordersku.update
//
// 只能更新发货前子订单的销售属性 &lt;br/&gt;只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 &lt;br/&gt;必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先
type TaobaoTradeOrderskuUpdateAPIRequest struct {
	model.Params
	// 销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
	_skuProps string
	// 子订单编号（对于单笔订单的交易可以传交易编号）。
	_oid int64
	// 销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
	_skuId int64
}

// NewTaobaoTradeOrderskuUpdateRequest 初始化TaobaoTradeOrderskuUpdateAPIRequest对象
func NewTaobaoTradeOrderskuUpdateRequest() *TaobaoTradeOrderskuUpdateAPIRequest {
	return &TaobaoTradeOrderskuUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeOrderskuUpdateAPIRequest) Reset() {
	r._skuProps = ""
	r._oid = 0
	r._skuId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeOrderskuUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.trade.ordersku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeOrderskuUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeOrderskuUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuProps is SkuProps Setter
// 销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
func (r *TaobaoTradeOrderskuUpdateAPIRequest) SetSkuProps(_skuProps string) error {
	r._skuProps = _skuProps
	r.Set("sku_props", _skuProps)
	return nil
}

// GetSkuProps SkuProps Getter
func (r TaobaoTradeOrderskuUpdateAPIRequest) GetSkuProps() string {
	return r._skuProps
}

// SetOid is Oid Setter
// 子订单编号（对于单笔订单的交易可以传交易编号）。
func (r *TaobaoTradeOrderskuUpdateAPIRequest) SetOid(_oid int64) error {
	r._oid = _oid
	r.Set("oid", _oid)
	return nil
}

// GetOid Oid Getter
func (r TaobaoTradeOrderskuUpdateAPIRequest) GetOid() int64 {
	return r._oid
}

// SetSkuId is SkuId Setter
// 销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
func (r *TaobaoTradeOrderskuUpdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoTradeOrderskuUpdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}

var poolTaobaoTradeOrderskuUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeOrderskuUpdateRequest()
	},
}

// GetTaobaoTradeOrderskuUpdateRequest 从 sync.Pool 获取 TaobaoTradeOrderskuUpdateAPIRequest
func GetTaobaoTradeOrderskuUpdateAPIRequest() *TaobaoTradeOrderskuUpdateAPIRequest {
	return poolTaobaoTradeOrderskuUpdateAPIRequest.Get().(*TaobaoTradeOrderskuUpdateAPIRequest)
}

// ReleaseTaobaoTradeOrderskuUpdateAPIRequest 将 TaobaoTradeOrderskuUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeOrderskuUpdateAPIRequest(v *TaobaoTradeOrderskuUpdateAPIRequest) {
	v.Reset()
	poolTaobaoTradeOrderskuUpdateAPIRequest.Put(v)
}
