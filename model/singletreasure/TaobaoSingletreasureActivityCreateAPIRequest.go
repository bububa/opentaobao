package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivitycreateAPIRequest 活动创建接口 API请求
// taobao.singletreasure.activity.create
//
// 创建优惠活动
type TaobaosingletreasureactivitycreateAPIRequest struct {
	model.Params
	// 系统入参
	_activityInfo *ActivityInfoCreateDto
}

// NewTaobaosingletreasureactivitycreateRequest 初始化TaobaosingletreasureactivitycreateAPIRequest对象
func NewTaobaosingletreasureactivitycreateRequest() *TaobaosingletreasureactivitycreateAPIRequest {
	return &TaobaosingletreasureactivitycreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosingletreasureactivitycreateAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosingletreasureactivitycreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosingletreasureactivitycreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityInfo is ActivityInfo Setter
// 系统入参
func (r *TaobaosingletreasureactivitycreateAPIRequest) SetActivityInfo(_activityInfo *ActivityInfoCreateDto) error {
	r._activityInfo = _activityInfo
	r.Set("activity_info", _activityInfo)
	return nil
}

// GetActivityInfo ActivityInfo Getter
func (r TaobaosingletreasureactivitycreateAPIRequest) GetActivityInfo() *ActivityInfoCreateDto {
	return r._activityInfo
}
