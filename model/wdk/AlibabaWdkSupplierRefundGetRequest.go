package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口按订单号批量查询供应商退款单 API请求
alibaba.wdk.supplier.refund.get

五道口按订单号批量查询供应商退款单
*/
type AlibabaWdkSupplierRefundGetRequest struct {
    model.Params
    // 查询入参
    _supplierRefundQueryListRequest   *SupplierRefundQueryListRequest
}

// 初始化AlibabaWdkSupplierRefundGetRequest对象
func NewAlibabaWdkSupplierRefundGetRequest() *AlibabaWdkSupplierRefundGetRequest{
    return &AlibabaWdkSupplierRefundGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierRefundGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.supplier.refund.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SupplierRefundQueryListRequest Setter
// 查询入参
func (r *AlibabaWdkSupplierRefundGetRequest) SetSupplierRefundQueryListRequest(_supplierRefundQueryListRequest *SupplierRefundQueryListRequest) error {
    r._supplierRefundQueryListRequest = _supplierRefundQueryListRequest
    r.Set("supplier_refund_query_list_request", _supplierRefundQueryListRequest)
    return nil
}

// SupplierRefundQueryListRequest Getter
func (r AlibabaWdkSupplierRefundGetRequest) GetSupplierRefundQueryListRequest() *SupplierRefundQueryListRequest {
    return r._supplierRefundQueryListRequest
}
