package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbInventoryDetailGetAPIRequest 查询库存详情 API请求
// taobao.wlb.inventory.detail.get
//
// 查询库存详情，通过商品ID获取发送请求的卖家的库存详情
type TaobaoWlbInventoryDetailGetAPIRequest struct {
	model.Params
	// 库存类型列表，值包括：<br/>VENDIBLE--可销售库存<br/>FREEZE--冻结库存<br/>ONWAY--在途库存<br/>DEFECT--残次品库存<br/>ENGINE_DAMAGE--机损<br/>BOX_DAMAGE--箱损<br/>EXPIRATION--过保
	_inventoryTypeList []string
	// 仓库编码
	_storeCode string
	// 商品ID
	_itemId int64
}

// NewTaobaoWlbInventoryDetailGetRequest 初始化TaobaoWlbInventoryDetailGetAPIRequest对象
func NewTaobaoWlbInventoryDetailGetRequest() *TaobaoWlbInventoryDetailGetAPIRequest {
	return &TaobaoWlbInventoryDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbInventoryDetailGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.inventory.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbInventoryDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInventoryTypeList is InventoryTypeList Setter
// 库存类型列表，值包括：<br/>VENDIBLE--可销售库存<br/>FREEZE--冻结库存<br/>ONWAY--在途库存<br/>DEFECT--残次品库存<br/>ENGINE_DAMAGE--机损<br/>BOX_DAMAGE--箱损<br/>EXPIRATION--过保
func (r *TaobaoWlbInventoryDetailGetAPIRequest) SetInventoryTypeList(_inventoryTypeList []string) error {
	r._inventoryTypeList = _inventoryTypeList
	r.Set("inventory_type_list", _inventoryTypeList)
	return nil
}

// GetInventoryTypeList InventoryTypeList Getter
func (r TaobaoWlbInventoryDetailGetAPIRequest) GetInventoryTypeList() []string {
	return r._inventoryTypeList
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaoWlbInventoryDetailGetAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoWlbInventoryDetailGetAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoWlbInventoryDetailGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoWlbInventoryDetailGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
