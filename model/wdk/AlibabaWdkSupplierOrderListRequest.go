package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口供应商维度正向订单拉取 API请求
alibaba.wdk.supplier.order.list

五道口供应商维度正向订单拉取
*/
type AlibabaWdkSupplierOrderListRequest struct {
    model.Params
    // 查询参数
    _supplierOrderQueryRequest   *SupplierOrderQueryRequest
}

// 初始化AlibabaWdkSupplierOrderListRequest对象
func NewAlibabaWdkSupplierOrderListRequest() *AlibabaWdkSupplierOrderListRequest{
    return &AlibabaWdkSupplierOrderListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierOrderListRequest) GetApiMethodName() string {
    return "alibaba.wdk.supplier.order.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierOrderListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SupplierOrderQueryRequest Setter
// 查询参数
func (r *AlibabaWdkSupplierOrderListRequest) SetSupplierOrderQueryRequest(_supplierOrderQueryRequest *SupplierOrderQueryRequest) error {
    r._supplierOrderQueryRequest = _supplierOrderQueryRequest
    r.Set("supplier_order_query_request", _supplierOrderQueryRequest)
    return nil
}

// SupplierOrderQueryRequest Getter
func (r AlibabaWdkSupplierOrderListRequest) GetSupplierOrderQueryRequest() *SupplierOrderQueryRequest {
    return r._supplierOrderQueryRequest
}
