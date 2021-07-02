package perfect

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPerfectPerformanceLocalitemPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.perfect.performance.localitem.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPerfectPerformanceLocalitemPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamPerfectPerformanceItemPublishReq Setter
// 请求参数
func (r *AlibabaPerfectPerformanceLocalitemPublishAPIRequest) SetParamPerfectPerformanceItemPublishReq(_paramPerfectPerformanceItemPublishReq *PerfectPerformanceItemPublishReq) error {
	r._paramPerfectPerformanceItemPublishReq = _paramPerfectPerformanceItemPublishReq
	r.Set("param_perfect_performance_item_publish_req", _paramPerfectPerformanceItemPublishReq)
	return nil
}

// Get ParamPerfectPerformanceItemPublishReq Getter
func (r AlibabaPerfectPerformanceLocalitemPublishAPIRequest) GetParamPerfectPerformanceItemPublishReq() *PerfectPerformanceItemPublishReq {
	return r._paramPerfectPerformanceItemPublishReq
}
