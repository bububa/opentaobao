package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodrugquantitybatchupdateAPIRequest 批量同步库存接口 API请求
// taobao.drug.quantity.batch.update
//
// 商家通过top接口可以批量修改商品库存
type TaobaodrugquantitybatchupdateAPIRequest struct {
	model.Params
	// 外部店铺ID
	_outStoreId string
	// 商品ID和库存
	_outItemIdQuantityMap string
}

// NewTaobaodrugquantitybatchupdateRequest 初始化TaobaodrugquantitybatchupdateAPIRequest对象
func NewTaobaodrugquantitybatchupdateRequest() *TaobaodrugquantitybatchupdateAPIRequest {
	return &TaobaodrugquantitybatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodrugquantitybatchupdateAPIRequest) GetApiMethodName() string {
	return "taobao.drug.quantity.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodrugquantitybatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodrugquantitybatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutStoreId is OutStoreId Setter
// 外部店铺ID
func (r *TaobaodrugquantitybatchupdateAPIRequest) SetOutStoreId(_outStoreId string) error {
	r._outStoreId = _outStoreId
	r.Set("out_store_id", _outStoreId)
	return nil
}

// GetOutStoreId OutStoreId Getter
func (r TaobaodrugquantitybatchupdateAPIRequest) GetOutStoreId() string {
	return r._outStoreId
}

// SetOutItemIdQuantityMap is OutItemIdQuantityMap Setter
// 商品ID和库存
func (r *TaobaodrugquantitybatchupdateAPIRequest) SetOutItemIdQuantityMap(_outItemIdQuantityMap string) error {
	r._outItemIdQuantityMap = _outItemIdQuantityMap
	r.Set("out_item_id_quantity_map", _outItemIdQuantityMap)
	return nil
}

// GetOutItemIdQuantityMap OutItemIdQuantityMap Getter
func (r TaobaodrugquantitybatchupdateAPIRequest) GetOutItemIdQuantityMap() string {
	return r._outItemIdQuantityMap
}
