package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemquantityupdateAPIRequest 宝贝/SKU库存修改 API请求
// taobao.item.quantity.update
//
// 提供按照全量或增量形式修改宝贝/SKU库存的功能
type TaobaoitemquantityupdateAPIRequest struct {
	model.Params
	// SKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU
	_outerId string
	// 商品数字ID，必填参数
	_numIid int64
	// 要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存
	_skuId int64
	// 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
	_quantity int64
	// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
	_type int64
}

// NewTaobaoitemquantityupdateRequest 初始化TaobaoitemquantityupdateAPIRequest对象
func NewTaobaoitemquantityupdateRequest() *TaobaoitemquantityupdateAPIRequest {
	return &TaobaoitemquantityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemquantityupdateAPIRequest) GetApiMethodName() string {
	return "taobao.item.quantity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemquantityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemquantityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// SKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU
func (r *TaobaoitemquantityupdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoitemquantityupdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetNumIid is NumIid Setter
// 商品数字ID，必填参数
func (r *TaobaoitemquantityupdateAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoitemquantityupdateAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetSkuId is SkuId Setter
// 要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存
func (r *TaobaoitemquantityupdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoitemquantityupdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetQuantity is Quantity Setter
// 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
func (r *TaobaoitemquantityupdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaoitemquantityupdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}

// SetType is Type Setter
// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
func (r *TaobaoitemquantityupdateAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoitemquantityupdateAPIRequest) GetType() int64 {
	return r._type
}
