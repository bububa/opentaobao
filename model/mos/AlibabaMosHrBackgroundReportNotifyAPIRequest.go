package mos

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
