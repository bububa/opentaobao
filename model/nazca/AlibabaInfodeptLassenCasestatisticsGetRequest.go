package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
法庭提交和结案案件量接口 API请求
alibaba.infodept.lassen.casestatistics.get

功能描述：云嘉为浙江省高院制作数据大屏，需展示网上法庭相关数据，我方为省高院提供浙江省内法院收案和结案的案件量，开放数据接口，供其调取这两组数据。
*/
type AlibabaInfodeptLassenCasestatisticsGetRequest struct {
    model.Params
    // 地区代码
    areaCode   string
    // 开始时间
    startTime   string
    // 结束时间
    endTime   string
}

// 初始化AlibabaInfodeptLassenCasestatisticsGetRequest对象
func NewAlibabaInfodeptLassenCasestatisticsGetRequest() *AlibabaInfodeptLassenCasestatisticsGetRequest{
    return &AlibabaInfodeptLassenCasestatisticsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInfodeptLassenCasestatisticsGetRequest) GetApiMethodName() string {
    return "alibaba.infodept.lassen.casestatistics.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInfodeptLassenCasestatisticsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AreaCode Setter
// 地区代码
func (r *AlibabaInfodeptLassenCasestatisticsGetRequest) SetAreaCode(areaCode string) error {
    r.areaCode = areaCode
    r.Set("area_code", areaCode)
    return nil
}

// AreaCode Getter
func (r AlibabaInfodeptLassenCasestatisticsGetRequest) GetAreaCode() string {
    return r.areaCode
}
// StartTime Setter
// 开始时间
func (r *AlibabaInfodeptLassenCasestatisticsGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlibabaInfodeptLassenCasestatisticsGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间
func (r *AlibabaInfodeptLassenCasestatisticsGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlibabaInfodeptLassenCasestatisticsGetRequest) GetEndTime() string {
    return r.endTime
}
