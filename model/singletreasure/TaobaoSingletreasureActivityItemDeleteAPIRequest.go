package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityItemDeleteAPIRequest 删除单品优惠接口 API请求
// taobao.singletreasure.activity.item.delete
//
// 删除单品优惠接口
type TaobaoSingletreasureActivityItemDeleteAPIRequest struct {
	model.Params
	// 活动id
	_activityId int64
	// 商品Id
	_itemId int64
}

// NewTaobaoSingletreasureActivityItemDeleteRequest 初始化TaobaoSingletreasureActivityItemDeleteAPIRequest对象
func NewTaobaoSingletreasureActivityItemDeleteRequest() *TaobaoSingletreasureActivityItemDeleteAPIRequest {
	return &TaobaoSingletreasureActivityItemDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// 活动id
func (r *TaobaoSingletreasureActivityItemDeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r TaobaoSingletreasureActivityItemDeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// Set is ItemId Setter
// 商品Id
func (r *TaobaoSingletreasureActivityItemDeleteAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoSingletreasureActivityItemDeleteAPIRequest) GetItemId() int64 {
	return r._itemId
}
