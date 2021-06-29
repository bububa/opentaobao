package mydata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取客户产品相关表现数据 APIRequest
alibaba.mydata.self.product.get

获取客户产品相关表现数据
*/
type AlibabaMydataSelfProductGetRequest struct {
    model.Params

    // 统计周期，可以为"day", "week", "month"
    statisticsType   string 

    // 统计日期
    statDate   string 

    // 待查询产品id列表
    productIds   []int64 

}

func NewAlibabaMydataSelfProductGetRequest() *AlibabaMydataSelfProductGetRequest{
    return &AlibabaMydataSelfProductGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMydataSelfProductGetRequest) GetApiMethodName() string {
    return "alibaba.mydata.self.product.get"
}

func (r AlibabaMydataSelfProductGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMydataSelfProductGetRequest) SetStatisticsType(statisticsType string) error {
    r.statisticsType = statisticsType
    r.Set("statistics_type", statisticsType)
    return nil
}

func (r AlibabaMydataSelfProductGetRequest) GetStatisticsType() string {
    return r.statisticsType
}

func (r *AlibabaMydataSelfProductGetRequest) SetStatDate(statDate string) error {
    r.statDate = statDate
    r.Set("stat_date", statDate)
    return nil
}

func (r AlibabaMydataSelfProductGetRequest) GetStatDate() string {
    return r.statDate
}

func (r *AlibabaMydataSelfProductGetRequest) SetProductIds(productIds []int64) error {
    r.productIds = productIds
    r.Set("product_ids", productIds)
    return nil
}

func (r AlibabaMydataSelfProductGetRequest) GetProductIds() []int64 {
    return r.productIds
}

