package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityupdateAPIRequest 修改活动接口 API请求
// taobao.singletreasure.activity.update
//
// 修改活动接口
type TaobaosingletreasureactivityupdateAPIRequest struct {
	model.Params
	// 系统入参
	_activityInfo *ActivityInfoCreateDto
}

// NewTaobaosingletreasureactivityupdateRequest 初始化TaobaosingletreasureactivityupdateAPIRequest对象
func NewTaobaosingletreasureactivityupdateRequest() *TaobaosingletreasureactivityupdateAPIRequest {
	return &TaobaosingletreasureactivityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosingletreasureactivityupdateAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosingletreasureactivityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosingletreasureactivityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityInfo is ActivityInfo Setter
// 系统入参
func (r *TaobaosingletreasureactivityupdateAPIRequest) SetActivityInfo(_activityInfo *ActivityInfoCreateDto) error {
	r._activityInfo = _activityInfo
	r.Set("activity_info", _activityInfo)
	return nil
}

// GetActivityInfo ActivityInfo Getter
func (r TaobaosingletreasureactivityupdateAPIRequest) GetActivityInfo() *ActivityInfoCreateDto {
	return r._activityInfo
}
