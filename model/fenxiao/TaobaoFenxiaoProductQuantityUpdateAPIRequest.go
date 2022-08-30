package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductQuantityUpdateAPIRequest 产品库存修改 API请求
// taobao.fenxiao.product.quantity.update
//
// 修改产品库存信息，支持全量修改以及增量修改两种方式
type TaobaoFenxiaoProductQuantityUpdateAPIRequest struct {
	model.Params
	// 库存修改值。产品有sku时，与sku属性顺序对应，用,分隔。产品无sku时，只写库存值。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
	_quantity string
	// sku属性值，产品有sku时填写，多个sku用,分隔。为空时默认该产品无sku，则只修改产品的库存。请参照taobao.fenxiao.products.get接口返回的properties设置入参
	_properties string
	// 产品ID
	_productId int64
	// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0
	_type int64
}

// NewTaobaoFenxiaoProductQuantityUpdateRequest 初始化TaobaoFenxiaoProductQuantityUpdateAPIRequest对象
func NewTaobaoFenxiaoProductQuantityUpdateRequest() *TaobaoFenxiaoProductQuantityUpdateAPIRequest {
	return &TaobaoFenxiaoProductQuantityUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductQuantityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.quantity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductQuantityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQuantity is Quantity Setter
// 库存修改值。产品有sku时，与sku属性顺序对应，用,分隔。产品无sku时，只写库存值。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
func (r *TaobaoFenxiaoProductQuantityUpdateAPIRequest) SetQuantity(_quantity string) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaoFenxiaoProductQuantityUpdateAPIRequest) GetQuantity() string {
	return r._quantity
}

// SetProperties is Properties Setter
// sku属性值，产品有sku时填写，多个sku用,分隔。为空时默认该产品无sku，则只修改产品的库存。请参照taobao.fenxiao.products.get接口返回的properties设置入参
func (r *TaobaoFenxiaoProductQuantityUpdateAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoFenxiaoProductQuantityUpdateAPIRequest) GetProperties() string {
	return r._properties
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductQuantityUpdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoFenxiaoProductQuantityUpdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetType is Type Setter
// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0
func (r *TaobaoFenxiaoProductQuantityUpdateAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoFenxiaoProductQuantityUpdateAPIRequest) GetType() int64 {
	return r._type
}
