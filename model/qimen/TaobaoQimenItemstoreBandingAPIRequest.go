package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenitemstorebandingAPIRequest 商品关联绑定接口 API请求
// taobao.qimen.itemstore.banding
//
// 商家在ERP等系统中调用该接口，将线上商品和线下门店“新建/删除”关联。这里的线上。每次只能单个商品关联多个门店，门店上限200
type TaobaoqimenitemstorebandingAPIRequest struct {
	model.Params
	// 门店列表
	_storeIds []int64
	// 备注信息
	_remark string
	// 操作类型
	_actionType string
	// 线上商品ID
	_itemId int64
}

// NewTaobaoqimenitemstorebandingRequest 初始化TaobaoqimenitemstorebandingAPIRequest对象
func NewTaobaoqimenitemstorebandingRequest() *TaobaoqimenitemstorebandingAPIRequest {
	return &TaobaoqimenitemstorebandingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenitemstorebandingAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.itemstore.banding"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenitemstorebandingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenitemstorebandingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreIds is StoreIds Setter
// 门店列表
func (r *TaobaoqimenitemstorebandingAPIRequest) SetStoreIds(_storeIds []int64) error {
	r._storeIds = _storeIds
	r.Set("store_ids", _storeIds)
	return nil
}

// GetStoreIds StoreIds Getter
func (r TaobaoqimenitemstorebandingAPIRequest) GetStoreIds() []int64 {
	return r._storeIds
}

// SetRemark is Remark Setter
// 备注信息
func (r *TaobaoqimenitemstorebandingAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoqimenitemstorebandingAPIRequest) GetRemark() string {
	return r._remark
}

// SetActionType is ActionType Setter
// 操作类型
func (r *TaobaoqimenitemstorebandingAPIRequest) SetActionType(_actionType string) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// GetActionType ActionType Getter
func (r TaobaoqimenitemstorebandingAPIRequest) GetActionType() string {
	return r._actionType
}

// SetItemId is ItemId Setter
// 线上商品ID
func (r *TaobaoqimenitemstorebandingAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoqimenitemstorebandingAPIRequest) GetItemId() int64 {
	return r._itemId
}
