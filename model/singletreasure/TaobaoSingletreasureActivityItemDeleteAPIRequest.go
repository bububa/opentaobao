package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityitemdeleteAPIRequest 删除单品优惠接口 API请求
// taobao.singletreasure.activity.item.delete
//
// 删除单品优惠接口
type TaobaosingletreasureactivityitemdeleteAPIRequest struct {
	model.Params
	// 活动id
	_activityId int64
	// 商品Id
	_itemId int64
}

// NewTaobaosingletreasureactivityitemdeleteRequest 初始化TaobaosingletreasureactivityitemdeleteAPIRequest对象
func NewTaobaosingletreasureactivityitemdeleteRequest() *TaobaosingletreasureactivityitemdeleteAPIRequest {
	return &TaobaosingletreasureactivityitemdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosingletreasureactivityitemdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosingletreasureactivityitemdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosingletreasureactivityitemdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *TaobaosingletreasureactivityitemdeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaosingletreasureactivityitemdeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetItemId is ItemId Setter
// 商品Id
func (r *TaobaosingletreasureactivityitemdeleteAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaosingletreasureactivityitemdeleteAPIRequest) GetItemId() int64 {
	return r._itemId
}
