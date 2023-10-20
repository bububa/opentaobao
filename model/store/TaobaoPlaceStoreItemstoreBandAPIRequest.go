package store

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreItemstoreBandAPIRequest 门店商品关联绑定接口 API请求
// taobao.place.store.itemstore.band
//
// 商品和多个门店关系绑定接口
type TaobaoPlaceStoreItemstoreBandAPIRequest struct {
	model.Params
	// 门店id
	_storeIds []string
	// 操作类型
	_actionType string
	// 商品id
	_itemId int64
}

// NewTaobaoPlaceStoreItemstoreBandRequest 初始化TaobaoPlaceStoreItemstoreBandAPIRequest对象
func NewTaobaoPlaceStoreItemstoreBandRequest() *TaobaoPlaceStoreItemstoreBandAPIRequest {
	return &TaobaoPlaceStoreItemstoreBandAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPlaceStoreItemstoreBandAPIRequest) Reset() {
	r._storeIds = r._storeIds[:0]
	r._actionType = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.itemstore.band"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreIds is StoreIds Setter
// 门店id
func (r *TaobaoPlaceStoreItemstoreBandAPIRequest) SetStoreIds(_storeIds []string) error {
	r._storeIds = _storeIds
	r.Set("store_ids", _storeIds)
	return nil
}

// GetStoreIds StoreIds Getter
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetStoreIds() []string {
	return r._storeIds
}

// SetActionType is ActionType Setter
// 操作类型
func (r *TaobaoPlaceStoreItemstoreBandAPIRequest) SetActionType(_actionType string) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// GetActionType ActionType Getter
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetActionType() string {
	return r._actionType
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoPlaceStoreItemstoreBandAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoPlaceStoreItemstoreBandAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPlaceStoreItemstoreBandRequest()
	},
}

// GetTaobaoPlaceStoreItemstoreBandRequest 从 sync.Pool 获取 TaobaoPlaceStoreItemstoreBandAPIRequest
func GetTaobaoPlaceStoreItemstoreBandAPIRequest() *TaobaoPlaceStoreItemstoreBandAPIRequest {
	return poolTaobaoPlaceStoreItemstoreBandAPIRequest.Get().(*TaobaoPlaceStoreItemstoreBandAPIRequest)
}

// ReleaseTaobaoPlaceStoreItemstoreBandAPIRequest 将 TaobaoPlaceStoreItemstoreBandAPIRequest 放入 sync.Pool
func ReleaseTaobaoPlaceStoreItemstoreBandAPIRequest(v *TaobaoPlaceStoreItemstoreBandAPIRequest) {
	v.Reset()
	poolTaobaoPlaceStoreItemstoreBandAPIRequest.Put(v)
}
