package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryInitialItemAPIRequest 商品库存初始化 API请求
// taobao.inventory.initial.item
//
// 建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
// 商家仓商品初始化在各个仓中库存
type TaobaoInventoryInitialItemAPIRequest struct {
	model.Params
	// 后端商品id
	_scItemId int64
	// 商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}]
	_storeInventorys string
}

// NewTaobaoInventoryInitialItemRequest 初始化TaobaoInventoryInitialItemAPIRequest对象
func NewTaobaoInventoryInitialItemRequest() *TaobaoInventoryInitialItemAPIRequest {
	return &TaobaoInventoryInitialItemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryInitialItemAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.initial.item"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryInitialItemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetScItemId is ScItemId Setter
// 后端商品id
func (r *TaobaoInventoryInitialItemAPIRequest) SetScItemId(_scItemId int64) error {
	r._scItemId = _scItemId
	r.Set("sc_item_id", _scItemId)
	return nil
}

// GetScItemId ScItemId Getter
func (r TaobaoInventoryInitialItemAPIRequest) GetScItemId() int64 {
	return r._scItemId
}

// SetStoreInventorys is StoreInventorys Setter
// 商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}]
func (r *TaobaoInventoryInitialItemAPIRequest) SetStoreInventorys(_storeInventorys string) error {
	r._storeInventorys = _storeInventorys
	r.Set("store_inventorys", _storeInventorys)
	return nil
}

// GetStoreInventorys StoreInventorys Getter
func (r TaobaoInventoryInitialItemAPIRequest) GetStoreInventorys() string {
	return r._storeInventorys
}
