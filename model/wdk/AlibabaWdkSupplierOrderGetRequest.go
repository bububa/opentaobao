package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口按订单号批量查询供应商正向订单 API请求
alibaba.wdk.supplier.order.get

五道口按订单号批量查询供应商正向订单
*/
type AlibabaWdkSupplierOrderGetRequest struct {
    model.Params
    // 查询参数
    supplierOrderQueryListRequest   *SupplierOrderQueryListRequest
}

// 初始化AlibabaWdkSupplierOrderGetRequest对象
func NewAlibabaWdkSupplierOrderGetRequest() *AlibabaWdkSupplierOrderGetRequest{
    return &AlibabaWdkSupplierOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierOrderGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.supplier.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SupplierOrderQueryListRequest Setter
// 查询参数
func (r *AlibabaWdkSupplierOrderGetRequest) SetSupplierOrderQueryListRequest(supplierOrderQueryListRequest *SupplierOrderQueryListRequest) error {
    r.supplierOrderQueryListRequest = supplierOrderQueryListRequest
    r.Set("supplier_order_query_list_request", supplierOrderQueryListRequest)
    return nil
}

// SupplierOrderQueryListRequest Getter
func (r AlibabaWdkSupplierOrderGetRequest) GetSupplierOrderQueryListRequest() *SupplierOrderQueryListRequest {
    return r.supplierOrderQueryListRequest
}
