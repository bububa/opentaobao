package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDrugQuantityUpdateAPIRequest 商家更新库存 API请求
// taobao.drug.quantity.update
//
// 商家通过top接口可以直接修改商品库存
type TaobaoDrugQuantityUpdateAPIRequest struct {
	model.Params
	// 外部店铺ID
	_outStoreId string
	// 外部商品ID
	_outItemId string
	// 库存数量
	_quantity int64
}

// NewTaobaoDrugQuantityUpdateRequest 初始化TaobaoDrugQuantityUpdateAPIRequest对象
func NewTaobaoDrugQuantityUpdateRequest() *TaobaoDrugQuantityUpdateAPIRequest {
	return &TaobaoDrugQuantityUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDrugQuantityUpdateAPIRequest) Reset() {
	r._outStoreId = ""
	r._outItemId = ""
	r._quantity = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDrugQuantityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.drug.quantity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDrugQuantityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDrugQuantityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutStoreId is OutStoreId Setter
// 外部店铺ID
func (r *TaobaoDrugQuantityUpdateAPIRequest) SetOutStoreId(_outStoreId string) error {
	r._outStoreId = _outStoreId
	r.Set("out_store_id", _outStoreId)
	return nil
}

// GetOutStoreId OutStoreId Getter
func (r TaobaoDrugQuantityUpdateAPIRequest) GetOutStoreId() string {
	return r._outStoreId
}

// SetOutItemId is OutItemId Setter
// 外部商品ID
func (r *TaobaoDrugQuantityUpdateAPIRequest) SetOutItemId(_outItemId string) error {
	r._outItemId = _outItemId
	r.Set("out_item_id", _outItemId)
	return nil
}

// GetOutItemId OutItemId Getter
func (r TaobaoDrugQuantityUpdateAPIRequest) GetOutItemId() string {
	return r._outItemId
}

// SetQuantity is Quantity Setter
// 库存数量
func (r *TaobaoDrugQuantityUpdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaoDrugQuantityUpdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}

var poolTaobaoDrugQuantityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDrugQuantityUpdateRequest()
	},
}

// GetTaobaoDrugQuantityUpdateRequest 从 sync.Pool 获取 TaobaoDrugQuantityUpdateAPIRequest
func GetTaobaoDrugQuantityUpdateAPIRequest() *TaobaoDrugQuantityUpdateAPIRequest {
	return poolTaobaoDrugQuantityUpdateAPIRequest.Get().(*TaobaoDrugQuantityUpdateAPIRequest)
}

// ReleaseTaobaoDrugQuantityUpdateAPIRequest 将 TaobaoDrugQuantityUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoDrugQuantityUpdateAPIRequest(v *TaobaoDrugQuantityUpdateAPIRequest) {
	v.Reset()
	poolTaobaoDrugQuantityUpdateAPIRequest.Put(v)
}
