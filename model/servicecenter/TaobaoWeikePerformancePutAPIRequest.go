package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweikeperformanceputAPIRequest 提交客服绩效接口 API请求
// taobao.weike.performance.put
//
// 提交客服绩效接口
type TaobaoweikeperformanceputAPIRequest struct {
	model.Params
	// 订单id
	_id int64
	// 绩效数据封装类
	_perInfoWrapper *PerformanceInfoWrapper
}

// NewTaobaoweikeperformanceputRequest 初始化TaobaoweikeperformanceputAPIRequest对象
func NewTaobaoweikeperformanceputRequest() *TaobaoweikeperformanceputAPIRequest {
	return &TaobaoweikeperformanceputAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoweikeperformanceputAPIRequest) GetApiMethodName() string {
	return "taobao.weike.performance.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoweikeperformanceputAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoweikeperformanceputAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 订单id
func (r *TaobaoweikeperformanceputAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoweikeperformanceputAPIRequest) GetId() int64 {
	return r._id
}

// SetPerInfoWrapper is PerInfoWrapper Setter
// 绩效数据封装类
func (r *TaobaoweikeperformanceputAPIRequest) SetPerInfoWrapper(_perInfoWrapper *PerformanceInfoWrapper) error {
	r._perInfoWrapper = _perInfoWrapper
	r.Set("per_info_wrapper", _perInfoWrapper)
	return nil
}

// GetPerInfoWrapper PerInfoWrapper Getter
func (r TaobaoweikeperformanceputAPIRequest) GetPerInfoWrapper() *PerformanceInfoWrapper {
	return r._perInfoWrapper
}
