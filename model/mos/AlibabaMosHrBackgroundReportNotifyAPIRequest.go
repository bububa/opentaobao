package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoshrbackgroundreportnotifyAPIRequest 背调公司背调结果通知 API请求
// alibaba.mos.hr.background.report.notify
//
// 背调公司背调结果通知
type AlibabamoshrbackgroundreportnotifyAPIRequest struct {
	model.Params
	// 背调结果通知对象
	_hrBackgroundReportNotifyDto *HrBackgroundReportNotifyDto
}

// NewAlibabamoshrbackgroundreportnotifyRequest 初始化AlibabamoshrbackgroundreportnotifyAPIRequest对象
func NewAlibabamoshrbackgroundreportnotifyRequest() *AlibabamoshrbackgroundreportnotifyAPIRequest {
	return &AlibabamoshrbackgroundreportnotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoshrbackgroundreportnotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.hr.background.report.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoshrbackgroundreportnotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoshrbackgroundreportnotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHrBackgroundReportNotifyDto is HrBackgroundReportNotifyDto Setter
// 背调结果通知对象
func (r *AlibabamoshrbackgroundreportnotifyAPIRequest) SetHrBackgroundReportNotifyDto(_hrBackgroundReportNotifyDto *HrBackgroundReportNotifyDto) error {
	r._hrBackgroundReportNotifyDto = _hrBackgroundReportNotifyDto
	r.Set("hr_background_report_notify_dto", _hrBackgroundReportNotifyDto)
	return nil
}

// GetHrBackgroundReportNotifyDto HrBackgroundReportNotifyDto Getter
func (r AlibabamoshrbackgroundreportnotifyAPIRequest) GetHrBackgroundReportNotifyDto() *HrBackgroundReportNotifyDto {
	return r._hrBackgroundReportNotifyDto
}
