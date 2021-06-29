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
    _pageNo   int64
    // 页大小，默认每页十条
    _pageSize   int64
    // 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
    _endPayTime   string
    // 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
    _startPayTime   string
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
func (r *AlibabaRetailCommissionOrderQueryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaRetailCommissionOrderQueryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 页大小，默认每页十条
func (r *AlibabaRetailCommissionOrderQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaRetailCommissionOrderQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// EndPayTime Setter
// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
func (r *AlibabaRetailCommissionOrderQueryRequest) SetEndPayTime(_endPayTime string) error {
    r._endPayTime = _endPayTime
    r.Set("end_pay_time", _endPayTime)
    return nil
}

// EndPayTime Getter
func (r AlibabaRetailCommissionOrderQueryRequest) GetEndPayTime() string {
    return r._endPayTime
}
// StartPayTime Setter
// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
func (r *AlibabaRetailCommissionOrderQueryRequest) SetStartPayTime(_startPayTime string) error {
    r._startPayTime = _startPayTime
    r.Set("start_pay_time", _startPayTime)
    return nil
}

// StartPayTime Getter
func (r AlibabaRetailCommissionOrderQueryRequest) GetStartPayTime() string {
    return r._startPayTime
}
