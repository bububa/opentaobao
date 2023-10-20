package singletreasure

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityUpdateAPIRequest 修改活动接口 API请求
// taobao.singletreasure.activity.update
//
// 修改活动接口
type TaobaoSingletreasureActivityUpdateAPIRequest struct {
	model.Params
	// 系统入参
	_activityInfo *ActivityInfoCreateDto
}

// NewTaobaoSingletreasureActivityUpdateRequest 初始化TaobaoSingletreasureActivityUpdateAPIRequest对象
func NewTaobaoSingletreasureActivityUpdateRequest() *TaobaoSingletreasureActivityUpdateAPIRequest {
	return &TaobaoSingletreasureActivityUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSingletreasureActivityUpdateAPIRequest) Reset() {
	r._activityInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSingletreasureActivityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityInfo is ActivityInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityUpdateAPIRequest) SetActivityInfo(_activityInfo *ActivityInfoCreateDto) error {
	r._activityInfo = _activityInfo
	r.Set("activity_info", _activityInfo)
	return nil
}

// GetActivityInfo ActivityInfo Getter
func (r TaobaoSingletreasureActivityUpdateAPIRequest) GetActivityInfo() *ActivityInfoCreateDto {
	return r._activityInfo
}

var poolTaobaoSingletreasureActivityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSingletreasureActivityUpdateRequest()
	},
}

// GetTaobaoSingletreasureActivityUpdateRequest 从 sync.Pool 获取 TaobaoSingletreasureActivityUpdateAPIRequest
func GetTaobaoSingletreasureActivityUpdateAPIRequest() *TaobaoSingletreasureActivityUpdateAPIRequest {
	return poolTaobaoSingletreasureActivityUpdateAPIRequest.Get().(*TaobaoSingletreasureActivityUpdateAPIRequest)
}

// ReleaseTaobaoSingletreasureActivityUpdateAPIRequest 将 TaobaoSingletreasureActivityUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSingletreasureActivityUpdateAPIRequest(v *TaobaoSingletreasureActivityUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSingletreasureActivityUpdateAPIRequest.Put(v)
}
