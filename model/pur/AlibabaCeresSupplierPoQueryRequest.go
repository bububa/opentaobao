package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
采购供应商订单查询接口 APIRequest
alibaba.ceres.supplier.po.query

采购供应商订单查询接口
*/
type AlibabaCeresSupplierPoQueryRequest struct {
    model.Params

    // 订单创建日期开始时间
    startDate   string 

    // 订单创建日期结束时间
    endDate   string 

    // 订单状态
    status   string 

}

func NewAlibabaCeresSupplierPoQueryRequest() *AlibabaCeresSupplierPoQueryRequest{
    return &AlibabaCeresSupplierPoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCeresSupplierPoQueryRequest) GetApiMethodName() string {
    return "alibaba.ceres.supplier.po.query"
}

func (r AlibabaCeresSupplierPoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCeresSupplierPoQueryRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r AlibabaCeresSupplierPoQueryRequest) GetStartDate() string {
    return r.startDate
}

func (r *AlibabaCeresSupplierPoQueryRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaCeresSupplierPoQueryRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaCeresSupplierPoQueryRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaCeresSupplierPoQueryRequest) GetStatus() string {
    return r.status
}

