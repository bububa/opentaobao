package perfect

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPerfectPerformanceLocalitemPublishAPIRequest 同城购定制化发品 API请求
// alibaba.perfect.performance.localitem.publish
//
// 同城购业务定制化发品接口，同城购业务线专用
type AlibabaPerfectPerformanceLocalitemPublishAPIRequest struct {
	model.Params
	// 请求参数
	_paramPerfectPerformanceItemPublishReq *PerfectPerformanceItemPublishReq
}

// NewAlibabaPerfectPerformanceLocalitemPublishRequest 初始化AlibabaPerfectPerformanceLocalitemPublishAPIRequest对象
func NewAlibabaPerfectPerformanceLocalitemPublishRequest() *AlibabaPerfectPerformanceLocalitemPublishAPIRequest {
	return &AlibabaPerfectPerformanceLocalitemPublishAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPerfectPerformanceLocalitemPublishAPIRequest) Reset() {
	r._paramPerfectPerformanceItemPublishReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPerfectPerformanceLocalitemPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.perfect.performance.localitem.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPerfectPerformanceLocalitemPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPerfectPerformanceLocalitemPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPerfectPerformanceItemPublishReq is ParamPerfectPerformanceItemPublishReq Setter
// 请求参数
func (r *AlibabaPerfectPerformanceLocalitemPublishAPIRequest) SetParamPerfectPerformanceItemPublishReq(_paramPerfectPerformanceItemPublishReq *PerfectPerformanceItemPublishReq) error {
	r._paramPerfectPerformanceItemPublishReq = _paramPerfectPerformanceItemPublishReq
	r.Set("param_perfect_performance_item_publish_req", _paramPerfectPerformanceItemPublishReq)
	return nil
}

// GetParamPerfectPerformanceItemPublishReq ParamPerfectPerformanceItemPublishReq Getter
func (r AlibabaPerfectPerformanceLocalitemPublishAPIRequest) GetParamPerfectPerformanceItemPublishReq() *PerfectPerformanceItemPublishReq {
	return r._paramPerfectPerformanceItemPublishReq
}

var poolAlibabaPerfectPerformanceLocalitemPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPerfectPerformanceLocalitemPublishRequest()
	},
}

// GetAlibabaPerfectPerformanceLocalitemPublishRequest 从 sync.Pool 获取 AlibabaPerfectPerformanceLocalitemPublishAPIRequest
func GetAlibabaPerfectPerformanceLocalitemPublishAPIRequest() *AlibabaPerfectPerformanceLocalitemPublishAPIRequest {
	return poolAlibabaPerfectPerformanceLocalitemPublishAPIRequest.Get().(*AlibabaPerfectPerformanceLocalitemPublishAPIRequest)
}

// ReleaseAlibabaPerfectPerformanceLocalitemPublishAPIRequest 将 AlibabaPerfectPerformanceLocalitemPublishAPIRequest 放入 sync.Pool
func ReleaseAlibabaPerfectPerformanceLocalitemPublishAPIRequest(v *AlibabaPerfectPerformanceLocalitemPublishAPIRequest) {
	v.Reset()
	poolAlibabaPerfectPerformanceLocalitemPublishAPIRequest.Put(v)
}
