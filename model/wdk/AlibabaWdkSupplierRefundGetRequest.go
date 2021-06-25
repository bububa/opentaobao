package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口按订单号批量查询供应商退款单 APIRequest
alibaba.wdk.supplier.refund.get

五道口按订单号批量查询供应商退款单
*/
type AlibabaWdkSupplierRefundGetRequest struct {
    model.Params

    // 查询入参
    supplierRefundQueryListRequest   *SupplierRefundQueryListRequest 

}

func NewAlibabaWdkSupplierRefundGetRequest() *AlibabaWdkSupplierRefundGetRequest{
    return &AlibabaWdkSupplierRefundGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSupplierRefundGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.supplier.refund.get"
}

func (r AlibabaWdkSupplierRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSupplierRefundGetRequest) SetSupplierRefundQueryListRequest(supplierRefundQueryListRequest *SupplierRefundQueryListRequest) error {
    r.supplierRefundQueryListRequest = supplierRefundQueryListRequest
    r.Set("supplier_refund_query_list_request", supplierRefundQueryListRequest)
    return nil
}

func (r AlibabaWdkSupplierRefundGetRequest) GetSupplierRefundQueryListRequest() *SupplierRefundQueryListRequest {
    return r.supplierRefundQueryListRequest
}

