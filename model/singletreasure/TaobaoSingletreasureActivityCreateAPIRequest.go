package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityCreateAPIRequest 活动创建接口 API请求
// taobao.singletreasure.activity.create
//
// 创建优惠活动
type TaobaoSingletreasureActivityCreateAPIRequest struct {
	model.Params
	// 系统入参
	_activityInfo *ActivityInfoCreateDto
}

// NewTaobaoSingletreasureActivityCreateRequest 初始化TaobaoSingletreasureActivityCreateAPIRequest对象
func NewTaobaoSingletreasureActivityCreateRequest() *TaobaoSingletreasureActivityCreateAPIRequest {
	return &TaobaoSingletreasureActivityCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityCreateAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityCreateAPIRequest) SetActivityInfo(_activityInfo *ActivityInfoCreateDto) error {
	r._activityInfo = _activityInfo
	r.Set("activity_info", _activityInfo)
	return nil
}

// Get ActivityInfo Getter
func (r TaobaoSingletreasureActivityCreateAPIRequest) GetActivityInfo() *ActivityInfoCreateDto {
	return r._activityInfo
}
