package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosHrBackgroundReportNotifyAPIRequest 背调公司背调结果通知 API请求
// alibaba.mos.hr.background.report.notify
//
// 背调公司背调结果通知
type AlibabaMosHrBackgroundReportNotifyAPIRequest struct {
	model.Params
	// 背调结果通知对象
	_hrBackgroundReportNotifyDto *HrBackgroundReportNotifyDto
}

// NewAlibabaMosHrBackgroundReportNotifyRequest 初始化AlibabaMosHrBackgroundReportNotifyAPIRequest对象
func NewAlibabaMosHrBackgroundReportNotifyRequest() *AlibabaMosHrBackgroundReportNotifyAPIRequest {
	return &AlibabaMosHrBackgroundReportNotifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosHrBackgroundReportNotifyAPIRequest) Reset() {
	r._hrBackgroundReportNotifyDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosHrBackgroundReportNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.hr.background.report.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosHrBackgroundReportNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosHrBackgroundReportNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHrBackgroundReportNotifyDto is HrBackgroundReportNotifyDto Setter
// 背调结果通知对象
func (r *AlibabaMosHrBackgroundReportNotifyAPIRequest) SetHrBackgroundReportNotifyDto(_hrBackgroundReportNotifyDto *HrBackgroundReportNotifyDto) error {
	r._hrBackgroundReportNotifyDto = _hrBackgroundReportNotifyDto
	r.Set("hr_background_report_notify_dto", _hrBackgroundReportNotifyDto)
	return nil
}

// GetHrBackgroundReportNotifyDto HrBackgroundReportNotifyDto Getter
func (r AlibabaMosHrBackgroundReportNotifyAPIRequest) GetHrBackgroundReportNotifyDto() *HrBackgroundReportNotifyDto {
	return r._hrBackgroundReportNotifyDto
}

var poolAlibabaMosHrBackgroundReportNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosHrBackgroundReportNotifyRequest()
	},
}

// GetAlibabaMosHrBackgroundReportNotifyRequest 从 sync.Pool 获取 AlibabaMosHrBackgroundReportNotifyAPIRequest
func GetAlibabaMosHrBackgroundReportNotifyAPIRequest() *AlibabaMosHrBackgroundReportNotifyAPIRequest {
	return poolAlibabaMosHrBackgroundReportNotifyAPIRequest.Get().(*AlibabaMosHrBackgroundReportNotifyAPIRequest)
}

// ReleaseAlibabaMosHrBackgroundReportNotifyAPIRequest 将 AlibabaMosHrBackgroundReportNotifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosHrBackgroundReportNotifyAPIRequest(v *AlibabaMosHrBackgroundReportNotifyAPIRequest) {
	v.Reset()
	poolAlibabaMosHrBackgroundReportNotifyAPIRequest.Put(v)
}
