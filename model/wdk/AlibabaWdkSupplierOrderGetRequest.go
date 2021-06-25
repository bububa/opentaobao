package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口按订单号批量查询供应商正向订单 APIRequest
alibaba.wdk.supplier.order.get

五道口按订单号批量查询供应商正向订单
*/
type AlibabaWdkSupplierOrderGetRequest struct {
    model.Params

    // 查询参数
    supplierOrderQueryListRequest   *SupplierOrderQueryListRequest 

}

func NewAlibabaWdkSupplierOrderGetRequest() *AlibabaWdkSupplierOrderGetRequest{
    return &AlibabaWdkSupplierOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSupplierOrderGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.supplier.order.get"
}

func (r AlibabaWdkSupplierOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSupplierOrderGetRequest) SetSupplierOrderQueryListRequest(supplierOrderQueryListRequest *SupplierOrderQueryListRequest) error {
    r.supplierOrderQueryListRequest = supplierOrderQueryListRequest
    r.Set("supplier_order_query_list_request", supplierOrderQueryListRequest)
    return nil
}

func (r AlibabaWdkSupplierOrderGetRequest) GetSupplierOrderQueryListRequest() *SupplierOrderQueryListRequest {
    return r.supplierOrderQueryListRequest
}

