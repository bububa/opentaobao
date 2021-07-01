package mydata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取客户产品相关表现数据 API请求
alibaba.mydata.self.product.get

获取客户产品相关表现数据
*/
type AlibabaMydataSelfProductGetAPIRequest struct {
    model.Params
    // 统计周期，可以为"day", "week", "month"
    _statisticsType   string
    // 统计日期
    _statDate   string
    // 待查询产品id列表
    _productIds   []int64
}

// 初始化AlibabaMydataSelfProductGetAPIRequest对象
func NewAlibabaMydataSelfProductGetRequest() *AlibabaMydataSelfProductGetAPIRequest{
    return &AlibabaMydataSelfProductGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMydataSelfProductGetAPIRequest) GetApiMethodName() string {
    return "alibaba.mydata.self.product.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMydataSelfProductGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StatisticsType Setter
// 统计周期，可以为"day", "week", "month"
func (r *AlibabaMydataSelfProductGetAPIRequest) SetStatisticsType(_statisticsType string) error {
    r._statisticsType = _statisticsType
    r.Set("statistics_type", _statisticsType)
    return nil
}

// StatisticsType Getter
func (r AlibabaMydataSelfProductGetAPIRequest) GetStatisticsType() string {
    return r._statisticsType
}
// StatDate Setter
// 统计日期
func (r *AlibabaMydataSelfProductGetAPIRequest) SetStatDate(_statDate string) error {
    r._statDate = _statDate
    r.Set("stat_date", _statDate)
    return nil
}

// StatDate Getter
func (r AlibabaMydataSelfProductGetAPIRequest) GetStatDate() string {
    return r._statDate
}
// ProductIds Setter
// 待查询产品id列表
func (r *AlibabaMydataSelfProductGetAPIRequest) SetProductIds(_productIds []int64) error {
    r._productIds = _productIds
    r.Set("product_ids", _productIds)
    return nil
}

// ProductIds Getter
func (r AlibabaMydataSelfProductGetAPIRequest) GetProductIds() []int64 {
    return r._productIds
}
