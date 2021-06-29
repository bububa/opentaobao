package mydata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取客户产品相关表现数据的可用时间范围 API请求
alibaba.mydata.self.product.date.get

获取客户产品相关表现数据的可用时间范围
*/
type AlibabaMydataSelfProductDateGetRequest struct {
    model.Params
    // 统计周期类型，可以为"day"，"week"，"month"
    statisticsType   string
}

// 初始化AlibabaMydataSelfProductDateGetRequest对象
func NewAlibabaMydataSelfProductDateGetRequest() *AlibabaMydataSelfProductDateGetRequest{
    return &AlibabaMydataSelfProductDateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMydataSelfProductDateGetRequest) GetApiMethodName() string {
    return "alibaba.mydata.self.product.date.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMydataSelfProductDateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StatisticsType Setter
// 统计周期类型，可以为"day"，"week"，"month"
func (r *AlibabaMydataSelfProductDateGetRequest) SetStatisticsType(statisticsType string) error {
    r.statisticsType = statisticsType
    r.Set("statistics_type", statisticsType)
    return nil
}

// StatisticsType Getter
func (r AlibabaMydataSelfProductDateGetRequest) GetStatisticsType() string {
    return r.statisticsType
}
