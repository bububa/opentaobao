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
    userId   int64
    // 报告类型
    reportType   string
    // 报告内容
    reportData   string
    // 报告周期
    reportPeriod   string
    // 报告时间
    reportTime   string
    // 报告周期天数
    reportPeriodDays   string
    // 数据来源
    reportSource   string
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
func (r *AlibabaFmhealthPressureReportCreateRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetUserId() int64 {
    return r.userId
}
// ReportType Setter
// 报告类型
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportType(reportType string) error {
    r.reportType = reportType
    r.Set("report_type", reportType)
    return nil
}

// ReportType Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportType() string {
    return r.reportType
}
// ReportData Setter
// 报告内容
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportData(reportData string) error {
    r.reportData = reportData
    r.Set("report_data", reportData)
    return nil
}

// ReportData Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportData() string {
    return r.reportData
}
// ReportPeriod Setter
// 报告周期
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportPeriod(reportPeriod string) error {
    r.reportPeriod = reportPeriod
    r.Set("report_period", reportPeriod)
    return nil
}

// ReportPeriod Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportPeriod() string {
    return r.reportPeriod
}
// ReportTime Setter
// 报告时间
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportTime(reportTime string) error {
    r.reportTime = reportTime
    r.Set("report_time", reportTime)
    return nil
}

// ReportTime Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportTime() string {
    return r.reportTime
}
// ReportPeriodDays Setter
// 报告周期天数
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportPeriodDays(reportPeriodDays string) error {
    r.reportPeriodDays = reportPeriodDays
    r.Set("report_period_days", reportPeriodDays)
    return nil
}

// ReportPeriodDays Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportPeriodDays() string {
    return r.reportPeriodDays
}
// ReportSource Setter
// 数据来源
func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportSource(reportSource string) error {
    r.reportSource = reportSource
    r.Set("report_source", reportSource)
    return nil
}

// ReportSource Getter
func (r AlibabaFmhealthPressureReportCreateRequest) GetReportSource() string {
    return r.reportSource
}
