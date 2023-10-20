package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaperfectperformancelocalitempublishAPIRequest 同城购定制化发品 API请求
// alibaba.perfect.performance.localitem.publish
//
// 同城购业务定制化发品接口，同城购业务线专用
type AlibabaperfectperformancelocalitempublishAPIRequest struct {
	model.Params
	// 请求参数
	_paramPerfectPerformanceItemPublishReq *PerfectPerformanceItemPublishReq
}

// NewAlibabaperfectperformancelocalitempublishRequest 初始化AlibabaperfectperformancelocalitempublishAPIRequest对象
func NewAlibabaperfectperformancelocalitempublishRequest() *AlibabaperfectperformancelocalitempublishAPIRequest {
	return &AlibabaperfectperformancelocalitempublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaperfectperformancelocalitempublishAPIRequest) GetApiMethodName() string {
	return "alibaba.perfect.performance.localitem.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaperfectperformancelocalitempublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaperfectperformancelocalitempublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPerfectPerformanceItemPublishReq is ParamPerfectPerformanceItemPublishReq Setter
// 请求参数
func (r *AlibabaperfectperformancelocalitempublishAPIRequest) SetParamPerfectPerformanceItemPublishReq(_paramPerfectPerformanceItemPublishReq *PerfectPerformanceItemPublishReq) error {
	r._paramPerfectPerformanceItemPublishReq = _paramPerfectPerformanceItemPublishReq
	r.Set("param_perfect_performance_item_publish_req", _paramPerfectPerformanceItemPublishReq)
	return nil
}

// GetParamPerfectPerformanceItemPublishReq ParamPerfectPerformanceItemPublishReq Getter
func (r AlibabaperfectperformancelocalitempublishAPIRequest) GetParamPerfectPerformanceItemPublishReq() *PerfectPerformanceItemPublishReq {
	return r._paramPerfectPerformanceItemPublishReq
}
