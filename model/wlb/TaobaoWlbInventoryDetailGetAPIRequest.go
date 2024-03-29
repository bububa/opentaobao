package wlb

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbInventoryDetailGetAPIRequest) Reset() {
	r._inventoryTypeList = r._inventoryTypeList[:0]
	r._storeCode = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbInventoryDetailGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.inventory.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbInventoryDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbInventoryDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryTypeList is InventoryTypeList Setter
// 库存类型列表，值包括：&lt;br/&gt;VENDIBLE--可销售库存&lt;br/&gt;FREEZE--冻结库存&lt;br/&gt;ONWAY--在途库存&lt;br/&gt;DEFECT--残次品库存&lt;br/&gt;ENGINE_DAMAGE--机损&lt;br/&gt;BOX_DAMAGE--箱损&lt;br/&gt;EXPIRATION--过保
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

var poolTaobaoWlbInventoryDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbInventoryDetailGetRequest()
	},
}

// GetTaobaoWlbInventoryDetailGetRequest 从 sync.Pool 获取 TaobaoWlbInventoryDetailGetAPIRequest
func GetTaobaoWlbInventoryDetailGetAPIRequest() *TaobaoWlbInventoryDetailGetAPIRequest {
	return poolTaobaoWlbInventoryDetailGetAPIRequest.Get().(*TaobaoWlbInventoryDetailGetAPIRequest)
}

// ReleaseTaobaoWlbInventoryDetailGetAPIRequest 将 TaobaoWlbInventoryDetailGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbInventoryDetailGetAPIRequest(v *TaobaoWlbInventoryDetailGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbInventoryDetailGetAPIRequest.Put(v)
}
