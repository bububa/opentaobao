package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest UD效果外投投放监测链接生成 API请求
// alibaba.taobao.ud.smart.monitor.url.create
//
// UD效果外投投放监测链接生成
type AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest struct {
	model.Params
	// 请求参数
	_monitorUrlTopRequestDto *MonitorUrlTopRequestDto
}

// NewAlibabaTaobaoUdSmartMonitorUrlCreateRequest 初始化AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest对象
func NewAlibabaTaobaoUdSmartMonitorUrlCreateRequest() *AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest {
	return &AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest) Reset() {
	r._monitorUrlTopRequestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.taobao.ud.smart.monitor.url.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMonitorUrlTopRequestDto is MonitorUrlTopRequestDto Setter
// 请求参数
func (r *AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest) SetMonitorUrlTopRequestDto(_monitorUrlTopRequestDto *MonitorUrlTopRequestDto) error {
	r._monitorUrlTopRequestDto = _monitorUrlTopRequestDto
	r.Set("monitor_url_top_request_dto", _monitorUrlTopRequestDto)
	return nil
}

// GetMonitorUrlTopRequestDto MonitorUrlTopRequestDto Getter
func (r AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest) GetMonitorUrlTopRequestDto() *MonitorUrlTopRequestDto {
	return r._monitorUrlTopRequestDto
}

var poolAlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTaobaoUdSmartMonitorUrlCreateRequest()
	},
}

// GetAlibabaTaobaoUdSmartMonitorUrlCreateRequest 从 sync.Pool 获取 AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest
func GetAlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest() *AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest {
	return poolAlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest.Get().(*AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest)
}

// ReleaseAlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest 将 AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest(v *AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest) {
	v.Reset()
	poolAlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest.Put(v)
}
