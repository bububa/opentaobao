package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
血压报告接口 API请求
alibaba.fmhealth.pressure.report.create

生成用户血压测量报告
*/
type AlibabaFmhealthPressureReportCreateRequest struct {
    model.Params
    // 用户id
    _userId   int64
    // 报告类型
    _reportType   string
    // 报告内容
    _reportData   string
    // 报告周期
    _reportPeriod   string
    // 报告时间
    _reportTime   string
    // 报告周期天数
    _reportPeriodDays   string
    // 数据来源
    _reportSource   string
}

// 初始化AlibabaFmhealthPressureReportCreateRequest对象
func NewAlibabaFmhealthPressureReportCreateRequest() *AlibabaFmhealthPressureReportCreateRequest{
    return &AlibabaFmhealthPressureReportCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFmhealthPressureReportCreateRequest) GetApiMethodName() string {
    return "alibaba.fmhealth.pressure.report.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFmhealthPressureReportCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id
func (r *AlibabaFmhealthPressureReportCreateRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetUserId() int64 {
    return r._userId
}
// ReportType Setter
// 报告类型
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportType(_reportType string) error {
    r._reportType = _reportType
    r.Set("report_type", _reportType)
    return nil
}

// ReportType Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportType() string {
    return r._reportType
}
// ReportData Setter
// 报告内容
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportData(_reportData string) error {
    r._reportData = _reportData
    r.Set("report_data", _reportData)
    return nil
}

// ReportData Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportData() string {
    return r._reportData
}
// ReportPeriod Setter
// 报告周期
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportPeriod(_reportPeriod string) error {
    r._reportPeriod = _reportPeriod
    r.Set("report_period", _reportPeriod)
    return nil
}

// ReportPeriod Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportPeriod() string {
    return r._reportPeriod
}
// ReportTime Setter
// 报告时间
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportTime(_reportTime string) error {
    r._reportTime = _reportTime
    r.Set("report_time", _reportTime)
    return nil
}

// ReportTime Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportTime() string {
    return r._reportTime
}
// ReportPeriodDays Setter
// 报告周期天数
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportPeriodDays(_reportPeriodDays string) error {
    r._reportPeriodDays = _reportPeriodDays
    r.Set("report_period_days", _reportPeriodDays)
    return nil
}

// ReportPeriodDays Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportPeriodDays() string {
    return r._reportPeriodDays
}
// ReportSource Setter
// 数据来源
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportSource(_reportSource string) error {
    r._reportSource = _reportSource
    r.Set("report_source", _reportSource)
    return nil
}

// ReportSource Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportSource() string {
    return r._reportSource
}