package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销订单查询 API请求
alibaba.retail.commission.order.query

查询商家的分销订单
*/
type AlibabaRetailCommissionOrderQueryRequest struct {
    model.Params
    // 页码，默认第一页
    pageNo   int64
    // 页大小，默认每页十条
    pageSize   int64
    // 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
    endPayTime   string
    // 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
    startPayTime   string
}

// 初始化AlibabaRetailCommissionOrderQueryRequest对象
func NewAlibabaRetailCommissionOrderQueryRequest() *AlibabaRetailCommissionOrderQueryRequest{
    return &AlibabaRetailCommissionOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailCommissionOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.retail.commission.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailCommissionOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNo Setter
// 页码，默认第一页
func (r *AlibabaRetailCommissionOrderQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaRetailCommissionOrderQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 页大小，默认每页十条
func (r *AlibabaRetailCommissionOrderQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaRetailCommissionOrderQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// EndPayTime Setter
// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
func (r *AlibabaRetailCommissionOrderQueryRequest) SetEndPayTime(endPayTime string) error {
    r.endPayTime = endPayTime
    r.Set("end_pay_time", endPayTime)
    return nil
}

// EndPayTime Getter
func (r AlibabaRetailCommissionOrderQueryRequest) GetEndPayTime() string {
    return r.endPayTime
}
// StartPayTime Setter
// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
func (r *AlibabaRetailCommissionOrderQueryRequest) SetStartPayTime(startPayTime string) error {
    r.startPayTime = startPayTime
    r.Set("start_pay_time", startPayTime)
    return nil
}

// StartPayTime Getter
func (r AlibabaRetailCommissionOrderQueryRequest) GetStartPayTime() string {
    return r.startPayTime
}
