package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodrugquantityupdateAPIRequest 商家更新库存 API请求
// taobao.drug.quantity.update
//
// 商家通过top接口可以直接修改商品库存
type TaobaodrugquantityupdateAPIRequest struct {
	model.Params
	// 外部店铺ID
	_outStoreId string
	// 外部商品ID
	_outItemId string
	// 库存数量
	_quantity int64
}

// NewTaobaodrugquantityupdateRequest 初始化TaobaodrugquantityupdateAPIRequest对象
func NewTaobaodrugquantityupdateRequest() *TaobaodrugquantityupdateAPIRequest {
	return &TaobaodrugquantityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodrugquantityupdateAPIRequest) GetApiMethodName() string {
	return "taobao.drug.quantity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodrugquantityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodrugquantityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutStoreId is OutStoreId Setter
// 外部店铺ID
func (r *TaobaodrugquantityupdateAPIRequest) SetOutStoreId(_outStoreId string) error {
	r._outStoreId = _outStoreId
	r.Set("out_store_id", _outStoreId)
	return nil
}

// GetOutStoreId OutStoreId Getter
func (r TaobaodrugquantityupdateAPIRequest) GetOutStoreId() string {
	return r._outStoreId
}

// SetOutItemId is OutItemId Setter
// 外部商品ID
func (r *TaobaodrugquantityupdateAPIRequest) SetOutItemId(_outItemId string) error {
	r._outItemId = _outItemId
	r.Set("out_item_id", _outItemId)
	return nil
}

// GetOutItemId OutItemId Getter
func (r TaobaodrugquantityupdateAPIRequest) GetOutItemId() string {
	return r._outItemId
}

// SetQuantity is Quantity Setter
// 库存数量
func (r *TaobaodrugquantityupdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaodrugquantityupdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}
