package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivitydeleteAPIRequest 删除活动接口 API请求
// taobao.singletreasure.activity.delete
//
// 删除优惠活动
type TaobaosingletreasureactivitydeleteAPIRequest struct {
	model.Params
	// 活动Id
	_activityId int64
}

// NewTaobaosingletreasureactivitydeleteRequest 初始化TaobaosingletreasureactivitydeleteAPIRequest对象
func NewTaobaosingletreasureactivitydeleteRequest() *TaobaosingletreasureactivitydeleteAPIRequest {
	return &TaobaosingletreasureactivitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosingletreasureactivitydeleteAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosingletreasureactivitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosingletreasureactivitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动Id
func (r *TaobaosingletreasureactivitydeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaosingletreasureactivitydeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}
