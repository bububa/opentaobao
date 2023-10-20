package singletreasure

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSingletreasureActivityCreateAPIRequest) Reset() {
	r._activityInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityCreateAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSingletreasureActivityCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityInfo is ActivityInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityCreateAPIRequest) SetActivityInfo(_activityInfo *ActivityInfoCreateDto) error {
	r._activityInfo = _activityInfo
	r.Set("activity_info", _activityInfo)
	return nil
}

// GetActivityInfo ActivityInfo Getter
func (r TaobaoSingletreasureActivityCreateAPIRequest) GetActivityInfo() *ActivityInfoCreateDto {
	return r._activityInfo
}

var poolTaobaoSingletreasureActivityCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSingletreasureActivityCreateRequest()
	},
}

// GetTaobaoSingletreasureActivityCreateRequest 从 sync.Pool 获取 TaobaoSingletreasureActivityCreateAPIRequest
func GetTaobaoSingletreasureActivityCreateAPIRequest() *TaobaoSingletreasureActivityCreateAPIRequest {
	return poolTaobaoSingletreasureActivityCreateAPIRequest.Get().(*TaobaoSingletreasureActivityCreateAPIRequest)
}

// ReleaseTaobaoSingletreasureActivityCreateAPIRequest 将 TaobaoSingletreasureActivityCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSingletreasureActivityCreateAPIRequest(v *TaobaoSingletreasureActivityCreateAPIRequest) {
	v.Reset()
	poolTaobaoSingletreasureActivityCreateAPIRequest.Put(v)
}
