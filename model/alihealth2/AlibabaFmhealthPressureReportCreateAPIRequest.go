package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFmhealthPressureReportCreateAPIRequest 血压报告接口 API请求
// alibaba.fmhealth.pressure.report.create
//
// 生成用户血压测量报告
type AlibabaFmhealthPressureReportCreateAPIRequest struct {
	model.Params
	// 报告类型
	_reportType string
	// 报告内容
	_reportData string
	// 报告周期
	_reportPeriod string
	// 报告时间
	_reportTime string
	// 报告周期天数
	_reportPeriodDays string
	// 数据来源
	_reportSource string
	// 用户id
	_userId int64
}

// NewAlibabaFmhealthPressureReportCreateRequest 初始化AlibabaFmhealthPressureReportCreateAPIRequest对象
func NewAlibabaFmhealthPressureReportCreateRequest() *AlibabaFmhealthPressureReportCreateAPIRequest {
	return &AlibabaFmhealthPressureReportCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFmhealthPressureReportCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.fmhealth.pressure.report.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFmhealthPressureReportCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReportType is ReportType Setter
// 报告类型
func (r *AlibabaFmhealthPressureReportCreateAPIRequest) SetReportType(_reportType string) error {
	r._reportType = _reportType
	r.Set("report_type", _reportType)
	return nil
}

// GetReportType ReportType Getter
func (r AlibabaFmhealthPressureReportCreateAPIRequest) GetReportType() string {
	return r._reportType
}

// SetReportData is ReportData Setter
// 报告内容
func (r *AlibabaFmhealthPressureReportCreateAPIRequest) SetReportData(_reportData string) error {
	r._reportData = _reportData
	r.Set("report_data", _reportData)
	return nil
}

// GetReportData ReportData Getter
func (r AlibabaFmhealthPressureReportCreateAPIRequest) GetReportData() string {
	return r._reportData
}

// SetReportPeriod is ReportPeriod Setter
// 报告周期
func (r *AlibabaFmhealthPressureReportCreateAPIRequest) SetReportPeriod(_reportPeriod string) error {
	r._reportPeriod = _reportPeriod
	r.Set("report_period", _reportPeriod)
	return nil
}

// GetReportPeriod ReportPeriod Getter
func (r AlibabaFmhealthPressureReportCreateAPIRequest) GetReportPeriod() string {
	return r._reportPeriod
}

// SetReportTime is ReportTime Setter
// 报告时间
func (r *AlibabaFmhealthPressureReportCreateAPIRequest) SetReportTime(_reportTime string) error {
	r._reportTime = _reportTime
	r.Set("report_time", _reportTime)
	return nil
}

// GetReportTime ReportTime Getter
func (r AlibabaFmhealthPressureReportCreateAPIRequest) GetReportTime() string {
	return r._reportTime
}

// SetReportPeriodDays is ReportPeriodDays Setter
// 报告周期天数
func (r *AlibabaFmhealthPressureReportCreateAPIRequest) SetReportPeriodDays(_reportPeriodDays string) error {
	r._reportPeriodDays = _reportPeriodDays
	r.Set("report_period_days", _reportPeriodDays)
	return nil
}

// GetReportPeriodDays ReportPeriodDays Getter
func (r AlibabaFmhealthPressureReportCreateAPIRequest) GetReportPeriodDays() string {
	return r._reportPeriodDays
}

// SetReportSource is ReportSource Setter
// 数据来源
func (r *AlibabaFmhealthPressureReportCreateAPIRequest) SetReportSource(_reportSource string) error {
	r._reportSource = _reportSource
	r.Set("report_source", _reportSource)
	return nil
}

// GetReportSource ReportSource Getter
func (r AlibabaFmhealthPressureReportCreateAPIRequest) GetReportSource() string {
	return r._reportSource
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaFmhealthPressureReportCreateAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaFmhealthPressureReportCreateAPIRequest) GetUserId() int64 {
	return r._userId
}
