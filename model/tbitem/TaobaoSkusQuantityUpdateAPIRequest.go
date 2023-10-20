package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSkusQuantityUpdateAPIRequest SKU库存修改 API请求
// taobao.skus.quantity.update
//
// 提供按照全量/增量的方式批量修改SKU库存的功能
type TaobaoSkusQuantityUpdateAPIRequest struct {
	model.Params
	// sku库存批量修改入参，用于指定一批sku和每个sku的库存修改值，特殊可填。格式为skuId:库存修改值;skuId:库存修改值。最多支持20个SKU同时修改。
	_skuidQuantities string
	// 特殊可选，skuIdQuantities为空的时候用该字段通过outerId来指定sku和其库存修改值。格式为outerId:库存修改值;outerId:库存修改值。当skuIdQuantities不为空的时候该字段失效。当一个outerId对应多个sku时，所有匹配到的sku都会被修改库存。最多支持20个SKU同时修改。
	_outeridQuantities string
	// 商品数字ID，必填参数
	_numIid int64
	// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0.
	_type int64
}

// NewTaobaoSkusQuantityUpdateRequest 初始化TaobaoSkusQuantityUpdateAPIRequest对象
func NewTaobaoSkusQuantityUpdateRequest() *TaobaoSkusQuantityUpdateAPIRequest {
	return &TaobaoSkusQuantityUpdateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSkusQuantityUpdateAPIRequest) Reset() {
	r._skuidQuantities = ""
	r._outeridQuantities = ""
	r._numIid = 0
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSkusQuantityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.skus.quantity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSkusQuantityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSkusQuantityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuidQuantities is SkuidQuantities Setter
// sku库存批量修改入参，用于指定一批sku和每个sku的库存修改值，特殊可填。格式为skuId:库存修改值;skuId:库存修改值。最多支持20个SKU同时修改。
func (r *TaobaoSkusQuantityUpdateAPIRequest) SetSkuidQuantities(_skuidQuantities string) error {
	r._skuidQuantities = _skuidQuantities
	r.Set("skuid_quantities", _skuidQuantities)
	return nil
}

// GetSkuidQuantities SkuidQuantities Getter
func (r TaobaoSkusQuantityUpdateAPIRequest) GetSkuidQuantities() string {
	return r._skuidQuantities
}

// SetOuteridQuantities is OuteridQuantities Setter
// 特殊可选，skuIdQuantities为空的时候用该字段通过outerId来指定sku和其库存修改值。格式为outerId:库存修改值;outerId:库存修改值。当skuIdQuantities不为空的时候该字段失效。当一个outerId对应多个sku时，所有匹配到的sku都会被修改库存。最多支持20个SKU同时修改。
func (r *TaobaoSkusQuantityUpdateAPIRequest) SetOuteridQuantities(_outeridQuantities string) error {
	r._outeridQuantities = _outeridQuantities
	r.Set("outerid_quantities", _outeridQuantities)
	return nil
}

// GetOuteridQuantities OuteridQuantities Getter
func (r TaobaoSkusQuantityUpdateAPIRequest) GetOuteridQuantities() string {
	return r._outeridQuantities
}

// SetNumIid is NumIid Setter
// 商品数字ID，必填参数
func (r *TaobaoSkusQuantityUpdateAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoSkusQuantityUpdateAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetType is Type Setter
// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0.
func (r *TaobaoSkusQuantityUpdateAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoSkusQuantityUpdateAPIRequest) GetType() int64 {
	return r._type
}

var poolTaobaoSkusQuantityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSkusQuantityUpdateRequest()
	},
}

// GetTaobaoSkusQuantityUpdateRequest 从 sync.Pool 获取 TaobaoSkusQuantityUpdateAPIRequest
func GetTaobaoSkusQuantityUpdateAPIRequest() *TaobaoSkusQuantityUpdateAPIRequest {
	return poolTaobaoSkusQuantityUpdateAPIRequest.Get().(*TaobaoSkusQuantityUpdateAPIRequest)
}

// ReleaseTaobaoSkusQuantityUpdateAPIRequest 将 TaobaoSkusQuantityUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSkusQuantityUpdateAPIRequest(v *TaobaoSkusQuantityUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSkusQuantityUpdateAPIRequest.Put(v)
}
