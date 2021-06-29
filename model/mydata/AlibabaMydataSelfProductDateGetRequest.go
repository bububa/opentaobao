package mydata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取客户产品相关表现数据的可用时间范围 APIRequest
alibaba.mydata.self.product.date.get

获取客户产品相关表现数据的可用时间范围
*/
type AlibabaMydataSelfProductDateGetRequest struct {
    model.Params

    // 统计周期类型，可以为"day"，"week"，"month"
    statisticsType   string 

}

func NewAlibabaMydataSelfProductDateGetRequest() *AlibabaMydataSelfProductDateGetRequest{
    return &AlibabaMydataSelfProductDateGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMydataSelfProductDateGetRequest) GetApiMethodName() string {
    return "alibaba.mydata.self.product.date.get"
}

func (r AlibabaMydataSelfProductDateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMydataSelfProductDateGetRequest) SetStatisticsType(statisticsType string) error {
    r.statisticsType = statisticsType
    r.Set("statistics_type", statisticsType)
    return nil
}

func (r AlibabaMydataSelfProductDateGetRequest) GetStatisticsType() string {
    return r.statisticsType
}

