package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityDeleteAPIRequest
删除活动接口 API请求
taobao.singletreasure.activity.delete

删除优惠活动 */
type TaobaoSingletreasureActivityDeleteAPIRequest struct {
	model.Params
	// 活动Id
	_activityId int64
}

// NewTaobaoSingletreasureActivityDeleteRequest 初始化TaobaoSingletreasureActivityDeleteAPIRequest对象
func NewTaobaoSingletreasureActivityDeleteRequest() *TaobaoSingletreasureActivityDeleteAPIRequest {
	return &TaobaoSingletreasureActivityDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// 活动Id
func (r *TaobaoSingletreasureActivityDeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r TaobaoSingletreasureActivityDeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}
