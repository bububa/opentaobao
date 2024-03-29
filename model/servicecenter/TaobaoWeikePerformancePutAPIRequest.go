package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikePerformancePutAPIRequest 提交客服绩效接口 API请求
// taobao.weike.performance.put
//
// 提交客服绩效接口
type TaobaoWeikePerformancePutAPIRequest struct {
	model.Params
	// 订单id
	_id int64
	// 绩效数据封装类
	_perInfoWrapper *PerformanceInfoWrapper
}

// NewTaobaoWeikePerformancePutRequest 初始化TaobaoWeikePerformancePutAPIRequest对象
func NewTaobaoWeikePerformancePutRequest() *TaobaoWeikePerformancePutAPIRequest {
	return &TaobaoWeikePerformancePutAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWeikePerformancePutAPIRequest) Reset() {
	r._id = 0
	r._perInfoWrapper = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeikePerformancePutAPIRequest) GetApiMethodName() string {
	return "taobao.weike.performance.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeikePerformancePutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWeikePerformancePutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 订单id
func (r *TaobaoWeikePerformancePutAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoWeikePerformancePutAPIRequest) GetId() int64 {
	return r._id
}

// SetPerInfoWrapper is PerInfoWrapper Setter
// 绩效数据封装类
func (r *TaobaoWeikePerformancePutAPIRequest) SetPerInfoWrapper(_perInfoWrapper *PerformanceInfoWrapper) error {
	r._perInfoWrapper = _perInfoWrapper
	r.Set("per_info_wrapper", _perInfoWrapper)
	return nil
}

// GetPerInfoWrapper PerInfoWrapper Getter
func (r TaobaoWeikePerformancePutAPIRequest) GetPerInfoWrapper() *PerformanceInfoWrapper {
	return r._perInfoWrapper
}

var poolTaobaoWeikePerformancePutAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWeikePerformancePutRequest()
	},
}

// GetTaobaoWeikePerformancePutRequest 从 sync.Pool 获取 TaobaoWeikePerformancePutAPIRequest
func GetTaobaoWeikePerformancePutAPIRequest() *TaobaoWeikePerformancePutAPIRequest {
	return poolTaobaoWeikePerformancePutAPIRequest.Get().(*TaobaoWeikePerformancePutAPIRequest)
}

// ReleaseTaobaoWeikePerformancePutAPIRequest 将 TaobaoWeikePerformancePutAPIRequest 放入 sync.Pool
func ReleaseTaobaoWeikePerformancePutAPIRequest(v *TaobaoWeikePerformancePutAPIRequest) {
	v.Reset()
	poolTaobaoWeikePerformancePutAPIRequest.Put(v)
}
