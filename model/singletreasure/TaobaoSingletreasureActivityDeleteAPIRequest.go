package singletreasure

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityDeleteAPIRequest 删除活动接口 API请求
// taobao.singletreasure.activity.delete
//
// 删除优惠活动
type TaobaoSingletreasureActivityDeleteAPIRequest struct {
	model.Params
	// 活动Id
	_activityId int64
}

// NewTaobaoSingletreasureActivityDeleteRequest 初始化TaobaoSingletreasureActivityDeleteAPIRequest对象
func NewTaobaoSingletreasureActivityDeleteRequest() *TaobaoSingletreasureActivityDeleteAPIRequest {
	return &TaobaoSingletreasureActivityDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSingletreasureActivityDeleteAPIRequest) Reset() {
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSingletreasureActivityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动Id
func (r *TaobaoSingletreasureActivityDeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoSingletreasureActivityDeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTaobaoSingletreasureActivityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSingletreasureActivityDeleteRequest()
	},
}

// GetTaobaoSingletreasureActivityDeleteRequest 从 sync.Pool 获取 TaobaoSingletreasureActivityDeleteAPIRequest
func GetTaobaoSingletreasureActivityDeleteAPIRequest() *TaobaoSingletreasureActivityDeleteAPIRequest {
	return poolTaobaoSingletreasureActivityDeleteAPIRequest.Get().(*TaobaoSingletreasureActivityDeleteAPIRequest)
}

// ReleaseTaobaoSingletreasureActivityDeleteAPIRequest 将 TaobaoSingletreasureActivityDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoSingletreasureActivityDeleteAPIRequest(v *TaobaoSingletreasureActivityDeleteAPIRequest) {
	v.Reset()
	poolTaobaoSingletreasureActivityDeleteAPIRequest.Put(v)
}
