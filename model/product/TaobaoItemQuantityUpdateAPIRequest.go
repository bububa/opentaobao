package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemQuantityUpdateAPIRequest
宝贝/SKU库存修改 API请求
taobao.item.quantity.update

提供按照全量或增量形式修改宝贝/SKU库存的功能 */
type TaobaoItemQuantityUpdateAPIRequest struct {
	model.Params
	// 商品数字ID，必填参数
	_numIid int64
	// 要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存
	_skuId int64
	// SKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU
	_outerId string
	// 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
	_quantity int64
	// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
	_type int64
}

// NewTaobaoItemQuantityUpdateRequest 初始化TaobaoItemQuantityUpdateAPIRequest对象
func NewTaobaoItemQuantityUpdateRequest() *TaobaoItemQuantityUpdateAPIRequest {
	return &TaobaoItemQuantityUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemQuantityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.item.quantity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemQuantityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NumIid Setter
// 商品数字ID，必填参数
func (r *TaobaoItemQuantityUpdateAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// Get NumIid Getter
func (r TaobaoItemQuantityUpdateAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// Set is SkuId Setter
// 要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存
func (r *TaobaoItemQuantityUpdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r TaobaoItemQuantityUpdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// Set is OuterId Setter
// SKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU
func (r *TaobaoItemQuantityUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TaobaoItemQuantityUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is Quantity Setter
// 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
func (r *TaobaoItemQuantityUpdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// Get Quantity Getter
func (r TaobaoItemQuantityUpdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}

// Set is Type Setter
// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
func (r *TaobaoItemQuantityUpdateAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoItemQuantityUpdateAPIRequest) GetType() int64 {
	return r._type
}
