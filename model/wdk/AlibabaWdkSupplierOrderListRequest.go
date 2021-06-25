package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口供应商维度正向订单拉取 APIRequest
alibaba.wdk.supplier.order.list

五道口供应商维度正向订单拉取
*/
type AlibabaWdkSupplierOrderListRequest struct {
    model.Params

    // 查询参数
    supplierOrderQueryRequest   *SupplierOrderQueryRequest 

}

func NewAlibabaWdkSupplierOrderListRequest() *AlibabaWdkSupplierOrderListRequest{
    return &AlibabaWdkSupplierOrderListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSupplierOrderListRequest) GetApiMethodName() string {
    return "alibaba.wdk.supplier.order.list"
}

func (r AlibabaWdkSupplierOrderListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSupplierOrderListRequest) SetSupplierOrderQueryRequest(supplierOrderQueryRequest *SupplierOrderQueryRequest) error {
    r.supplierOrderQueryRequest = supplierOrderQueryRequest
    r.Set("supplier_order_query_request", supplierOrderQueryRequest)
    return nil
}

func (r AlibabaWdkSupplierOrderListRequest) GetSupplierOrderQueryRequest() *SupplierOrderQueryRequest {
    return r.supplierOrderQueryRequest
}

