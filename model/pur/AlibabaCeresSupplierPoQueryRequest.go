package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
采购供应商订单查询接口 API请求
alibaba.ceres.supplier.po.query

采购供应商订单查询接口
*/
type AlibabaCeresSupplierPoQueryRequest struct {
    model.Params
    // 订单创建日期开始时间
    _startDate   string
    // 订单创建日期结束时间
    _endDate   string
    // 订单状态
    _status   string
}

// 初始化AlibabaCeresSupplierPoQueryRequest对象
func NewAlibabaCeresSupplierPoQueryRequest() *AlibabaCeresSupplierPoQueryRequest{
    return &AlibabaCeresSupplierPoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCeresSupplierPoQueryRequest) GetApiMethodName() string {
    return "alibaba.ceres.supplier.po.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCeresSupplierPoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 订单创建日期开始时间
func (r *AlibabaCeresSupplierPoQueryRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r AlibabaCeresSupplierPoQueryRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 订单创建日期结束时间
func (r *AlibabaCeresSupplierPoQueryRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaCeresSupplierPoQueryRequest) GetEndDate() string {
    return r._endDate
}
// Status Setter
// 订单状态
func (r *AlibabaCeresSupplierPoQueryRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaCeresSupplierPoQueryRequest) GetStatus() string {
    return r._status
}
