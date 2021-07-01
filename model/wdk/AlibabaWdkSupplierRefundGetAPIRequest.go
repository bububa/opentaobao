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
type AlibabaWdkSupplierRefundGetAPIRequest struct {
    model.Params
    // 查询入参
    _supplierRefundQueryListRequest   *SupplierRefundQueryListRequest
}

// 初始化AlibabaWdkSupplierRefundGetAPIRequest对象
func NewAlibabaWdkSupplierRefundGetRequest() *AlibabaWdkSupplierRefundGetAPIRequest{
    return &AlibabaWdkSupplierRefundGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierRefundGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.supplier.refund.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierRefundGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SupplierRefundQueryListRequest Setter
// 查询入参
func (r *AlibabaWdkSupplierRefundGetAPIRequest) SetSupplierRefundQueryListRequest(_supplierRefundQueryListRequest *SupplierRefundQueryListRequest) error {
    r._supplierRefundQueryListRequest = _supplierRefundQueryListRequest
    r.Set("supplier_refund_query_list_request", _supplierRefundQueryListRequest)
    return nil
}

// SupplierRefundQueryListRequest Getter
func (r AlibabaWdkSupplierRefundGetAPIRequest) GetSupplierRefundQueryListRequest() *SupplierRefundQueryListRequest {
    return r._supplierRefundQueryListRequest
}
