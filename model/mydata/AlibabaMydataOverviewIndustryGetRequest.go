package mydata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的效果-获取Top行业列表 APIRequest
alibaba.mydata.overview.industry.get

获取数据管家我的效果API可以使用的行业
*/
type AlibabaMydataOverviewIndustryGetRequest struct {
    model.Params

    // 系统自动生成
    dateRange   *DateRange 

}

func NewAlibabaMydataOverviewIndustryGetRequest() *AlibabaMydataOverviewIndustryGetRequest{
    return &AlibabaMydataOverviewIndustryGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMydataOverviewIndustryGetRequest) GetApiMethodName() string {
    return "alibaba.mydata.overview.industry.get"
}

func (r AlibabaMydataOverviewIndustryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMydataOverviewIndustryGetRequest) SetDateRange(dateRange *DateRange) error {
    r.dateRange = dateRange
    r.Set("date_range", dateRange)
    return nil
}

func (r AlibabaMydataOverviewIndustryGetRequest) GetDateRange() *DateRange {
    return r.dateRange
}

