package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPerfectPerformanceLocalitemEditAPIRequest 同城购定制发品编辑 API请求
// alibaba.perfect.performance.localitem.edit
//
// 同城购业务定制化发品接口，同城购业务线专用
type AlibabaPerfectPerformanceLocalitemEditAPIRequest struct {
	model.Params
	// 请求参数
	_paramPerfectPerformanceItemPublishReq *PerfectPerformanceItemPublishReq
}

// NewAlibabaPerfectPerformanceLocalitemEditRequest 初始化AlibabaPerfectPerformanceLocalitemEditAPIRequest对象
func NewAlibabaPerfectPerformanceLocalitemEditRequest() *AlibabaPerfectPerformanceLocalitemEditAPIRequest {
	return &AlibabaPerfectPerformanceLocalitemEditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPerfectPerformanceLocalitemEditAPIRequest) GetApiMethodName() string {
	return "alibaba.perfect.performance.localitem.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPerfectPerformanceLocalitemEditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamPerfectPerformanceItemPublishReq Setter
// 请求参数
func (r *AlibabaPerfectPerformanceLocalitemEditAPIRequest) SetParamPerfectPerformanceItemPublishReq(_paramPerfectPerformanceItemPublishReq *PerfectPerformanceItemPublishReq) error {
	r._paramPerfectPerformanceItemPublishReq = _paramPerfectPerformanceItemPublishReq
	r.Set("param_perfect_performance_item_publish_req", _paramPerfectPerformanceItemPublishReq)
	return nil
}

// Get ParamPerfectPerformanceItemPublishReq Getter
func (r AlibabaPerfectPerformanceLocalitemEditAPIRequest) GetParamPerfectPerformanceItemPublishReq() *PerfectPerformanceItemPublishReq {
	return r._paramPerfectPerformanceItemPublishReq
}
