package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口按供应商拉取退款单 API请求
alibaba.wdk.supplier.refund.list

五道口按供应商拉取退款单
*/
type AlibabaWdkSupplierRefundListRequest struct {
    model.Params
    // 查询参数
    supplierRefundQueryRequest   *SupplierRefundQueryRequest
}

// 初始化AlibabaWdkSupplierRefundListRequest对象
func NewAlibabaWdkSupplierRefundListRequest() *AlibabaWdkSupplierRefundListRequest{
    return &AlibabaWdkSupplierRefundListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierRefundListRequest) GetApiMethodName() string {
    return "alibaba.wdk.supplier.refund.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierRefundListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SupplierRefundQueryRequest Setter
// 查询参数
func (r *AlibabaWdkSupplierRefundListRequest) SetSupplierRefundQueryRequest(supplierRefundQueryRequest *SupplierRefundQueryRequest) error {
    r.supplierRefundQueryRequest = supplierRefundQueryRequest
    r.Set("supplier_refund_query_request", supplierRefundQueryRequest)
    return nil
}

// SupplierRefundQueryRequest Getter
func (r AlibabaWdkSupplierRefundListRequest) GetSupplierRefundQueryRequest() *SupplierRefundQueryRequest {
    return r.supplierRefundQueryRequest
}
