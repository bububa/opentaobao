package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestoreitemstorebandAPIRequest 门店商品关联绑定接口 API请求
// taobao.place.store.itemstore.band
//
// 商品和多个门店关系绑定接口
type TaobaoplacestoreitemstorebandAPIRequest struct {
	model.Params
	// 门店id
	_storeIds []string
	// 操作类型
	_actionType string
	// 商品id
	_itemId int64
}

// NewTaobaoplacestoreitemstorebandRequest 初始化TaobaoplacestoreitemstorebandAPIRequest对象
func NewTaobaoplacestoreitemstorebandRequest() *TaobaoplacestoreitemstorebandAPIRequest {
	return &TaobaoplacestoreitemstorebandAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestoreitemstorebandAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.itemstore.band"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestoreitemstorebandAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestoreitemstorebandAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreIds is StoreIds Setter
// 门店id
func (r *TaobaoplacestoreitemstorebandAPIRequest) SetStoreIds(_storeIds []string) error {
	r._storeIds = _storeIds
	r.Set("store_ids", _storeIds)
	return nil
}

// GetStoreIds StoreIds Getter
func (r TaobaoplacestoreitemstorebandAPIRequest) GetStoreIds() []string {
	return r._storeIds
}

// SetActionType is ActionType Setter
// 操作类型
func (r *TaobaoplacestoreitemstorebandAPIRequest) SetActionType(_actionType string) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// GetActionType ActionType Getter
func (r TaobaoplacestoreitemstorebandAPIRequest) GetActionType() string {
	return r._actionType
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoplacestoreitemstorebandAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoplacestoreitemstorebandAPIRequest) GetItemId() int64 {
	return r._itemId
}
