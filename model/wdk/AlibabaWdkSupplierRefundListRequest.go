package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口按供应商拉取退款单 APIRequest
alibaba.wdk.supplier.refund.list

五道口按供应商拉取退款单
*/
type AlibabaWdkSupplierRefundListRequest struct {
    model.Params

    // 查询参数
    supplierRefundQueryRequest   *SupplierRefundQueryRequest 

}

func NewAlibabaWdkSupplierRefundListRequest() *AlibabaWdkSupplierRefundListRequest{
    return &AlibabaWdkSupplierRefundListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSupplierRefundListRequest) GetApiMethodName() string {
    return "alibaba.wdk.supplier.refund.list"
}

func (r AlibabaWdkSupplierRefundListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSupplierRefundListRequest) SetSupplierRefundQueryRequest(supplierRefundQueryRequest *SupplierRefundQueryRequest) error {
    r.supplierRefundQueryRequest = supplierRefundQueryRequest
    r.Set("supplier_refund_query_request", supplierRefundQueryRequest)
    return nil
}

func (r AlibabaWdkSupplierRefundListRequest) GetSupplierRefundQueryRequest() *SupplierRefundQueryRequest {
    return r.supplierRefundQueryRequest
}

