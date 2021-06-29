package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
血压报告接口 APIRequest
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

func NewAlibabaFmhealthPressureReportCreateRequest() *AlibabaFmhealthPressureReportCreateRequest{
    return &AlibabaFmhealthPressureReportCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFmhealthPressureReportCreateRequest) GetApiMethodName() string {
    return "alibaba.fmhealth.pressure.report.create"
}

func (r AlibabaFmhealthPressureReportCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFmhealthPressureReportCreateRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaFmhealthPressureReportCreateRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportType(reportType string) error {
    r.reportType = reportType
    r.Set("report_type", reportType)
    return nil
}

func (r AlibabaFmhealthPressureReportCreateRequest) GetReportType() string {
    return r.reportType
}

func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportData(reportData string) error {
    r.reportData = reportData
    r.Set("report_data", reportData)
    return nil
}

func (r AlibabaFmhealthPressureReportCreateRequest) GetReportData() string {
    return r.reportData
}

func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportPeriod(reportPeriod string) error {
    r.reportPeriod = reportPeriod
    r.Set("report_period", reportPeriod)
    return nil
}

func (r AlibabaFmhealthPressureReportCreateRequest) GetReportPeriod() string {
    return r.reportPeriod
}

func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportTime(reportTime string) error {
    r.reportTime = reportTime
    r.Set("report_time", reportTime)
    return nil
}

func (r AlibabaFmhealthPressureReportCreateRequest) GetReportTime() string {
    return r.reportTime
}

func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportPeriodDays(reportPeriodDays string) error {
    r.reportPeriodDays = reportPeriodDays
    r.Set("report_period_days", reportPeriodDays)
    return nil
}

func (r AlibabaFmhealthPressureReportCreateRequest) GetReportPeriodDays() string {
    return r.reportPeriodDays
}

func (r *AlibabaFmhealthPressureReportCreateRequest) SetReportSource(reportSource string) error {
    r.reportSource = reportSource
    r.Set("report_source", reportSource)
    return nil
}

func (r AlibabaFmhealthPressureReportCreateRequest) GetReportSource() string {
    return r.reportSource
}

