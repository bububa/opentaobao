package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDrugQuantityUpdateAPIRequest
商家更新库存 API请求
taobao.drug.quantity.update

商家通过top接口可以直接修改商品库存 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDrugQuantityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.drug.quantity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDrugQuantityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OutStoreId Setter
// 外部店铺ID
func (r *TaobaoDrugQuantityUpdateAPIRequest) SetOutStoreId(_outStoreId string) error {
	r._outStoreId = _outStoreId
	r.Set("out_store_id", _outStoreId)
	return nil
}

// Get OutStoreId Getter
func (r TaobaoDrugQuantityUpdateAPIRequest) GetOutStoreId() string {
	return r._outStoreId
}

// Set is OutItemId Setter
// 外部商品ID
func (r *TaobaoDrugQuantityUpdateAPIRequest) SetOutItemId(_outItemId string) error {
	r._outItemId = _outItemId
	r.Set("out_item_id", _outItemId)
	return nil
}

// Get OutItemId Getter
func (r TaobaoDrugQuantityUpdateAPIRequest) GetOutItemId() string {
	return r._outItemId
}

// Set is Quantity Setter
// 库存数量
func (r *TaobaoDrugQuantityUpdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// Get Quantity Getter
func (r TaobaoDrugQuantityUpdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}
