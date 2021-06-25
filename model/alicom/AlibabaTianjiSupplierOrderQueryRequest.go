package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询供应商订单 APIRequest
alibaba.tianji.supplier.order.query

查询供应商订单
*/
type AlibabaTianjiSupplierOrderQueryRequest struct {
    model.Params

    // 订单查询入参
    paramSupplierTopQueryModel   *SupplierTopQueryModel 

}

func NewAlibabaTianjiSupplierOrderQueryRequest() *AlibabaTianjiSupplierOrderQueryRequest{
    return &AlibabaTianjiSupplierOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTianjiSupplierOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.tianji.supplier.order.query"
}

func (r AlibabaTianjiSupplierOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTianjiSupplierOrderQueryRequest) SetParamSupplierTopQueryModel(paramSupplierTopQueryModel *SupplierTopQueryModel) error {
    r.paramSupplierTopQueryModel = paramSupplierTopQueryModel
    r.Set("param_supplier_top_query_model", paramSupplierTopQueryModel)
    return nil
}

func (r AlibabaTianjiSupplierOrderQueryRequest) GetParamSupplierTopQueryModel() *SupplierTopQueryModel {
    return r.paramSupplierTopQueryModel
}

