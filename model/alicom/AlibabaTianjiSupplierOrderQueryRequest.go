package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询供应商订单 API请求
alibaba.tianji.supplier.order.query

查询供应商订单
*/
type AlibabaTianjiSupplierOrderQueryRequest struct {
    model.Params
    // 订单查询入参
    _paramSupplierTopQueryModel   *SupplierTopQueryModel
}

// 初始化AlibabaTianjiSupplierOrderQueryRequest对象
func NewAlibabaTianjiSupplierOrderQueryRequest() *AlibabaTianjiSupplierOrderQueryRequest{
    return &AlibabaTianjiSupplierOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTianjiSupplierOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.tianji.supplier.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTianjiSupplierOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamSupplierTopQueryModel Setter
// 订单查询入参
func (r *AlibabaTianjiSupplierOrderQueryRequest) SetParamSupplierTopQueryModel(_paramSupplierTopQueryModel *SupplierTopQueryModel) error {
    r._paramSupplierTopQueryModel = _paramSupplierTopQueryModel
    r.Set("param_supplier_top_query_model", _paramSupplierTopQueryModel)
    return nil
}

// ParamSupplierTopQueryModel Getter
func (r AlibabaTianjiSupplierOrderQueryRequest) GetParamSupplierTopQueryModel() *SupplierTopQueryModel {
    return r._paramSupplierTopQueryModel
}
