package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaelophyorderlogisticstracecallbackAPIRequest 配送轨迹回传 API请求
// alibaba.aelophy.order.logistics.trace.callback
//
// 配送轨迹回传
type AlibabaaelophyorderlogisticstracecallbackAPIRequest struct {
	model.Params
	// 配送轨迹回传请求
	_logisticsTraceCallbackRequest *LogisticsTraceCallbackRequest
}

// NewAlibabaaelophyorderlogisticstracecallbackRequest 初始化AlibabaaelophyorderlogisticstracecallbackAPIRequest对象
func NewAlibabaaelophyorderlogisticstracecallbackRequest() *AlibabaaelophyorderlogisticstracecallbackAPIRequest {
	return &AlibabaaelophyorderlogisticstracecallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaelophyorderlogisticstracecallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.order.logistics.trace.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaelophyorderlogisticstracecallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaelophyorderlogisticstracecallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsTraceCallbackRequest is LogisticsTraceCallbackRequest Setter
// 配送轨迹回传请求
func (r *AlibabaaelophyorderlogisticstracecallbackAPIRequest) SetLogisticsTraceCallbackRequest(_logisticsTraceCallbackRequest *LogisticsTraceCallbackRequest) error {
	r._logisticsTraceCallbackRequest = _logisticsTraceCallbackRequest
	r.Set("logistics_trace_callback_request", _logisticsTraceCallbackRequest)
	return nil
}

// GetLogisticsTraceCallbackRequest LogisticsTraceCallbackRequest Getter
func (r AlibabaaelophyorderlogisticstracecallbackAPIRequest) GetLogisticsTraceCallbackRequest() *LogisticsTraceCallbackRequest {
	return r._logisticsTraceCallbackRequest
}
