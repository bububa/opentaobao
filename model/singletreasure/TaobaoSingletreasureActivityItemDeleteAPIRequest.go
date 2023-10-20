package singletreasure

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSingletreasureActivityItemDeleteAPIRequest) Reset() {
	r._activityId = 0
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSingletreasureActivityItemDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *TaobaoSingletreasureActivityItemDeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoSingletreasureActivityItemDeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetItemId is ItemId Setter
// 商品Id
func (r *TaobaoSingletreasureActivityItemDeleteAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoSingletreasureActivityItemDeleteAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoSingletreasureActivityItemDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSingletreasureActivityItemDeleteRequest()
	},
}

// GetTaobaoSingletreasureActivityItemDeleteRequest 从 sync.Pool 获取 TaobaoSingletreasureActivityItemDeleteAPIRequest
func GetTaobaoSingletreasureActivityItemDeleteAPIRequest() *TaobaoSingletreasureActivityItemDeleteAPIRequest {
	return poolTaobaoSingletreasureActivityItemDeleteAPIRequest.Get().(*TaobaoSingletreasureActivityItemDeleteAPIRequest)
}

// ReleaseTaobaoSingletreasureActivityItemDeleteAPIRequest 将 TaobaoSingletreasureActivityItemDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoSingletreasureActivityItemDeleteAPIRequest(v *TaobaoSingletreasureActivityItemDeleteAPIRequest) {
	v.Reset()
	poolTaobaoSingletreasureActivityItemDeleteAPIRequest.Put(v)
}
