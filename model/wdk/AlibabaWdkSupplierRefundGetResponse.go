package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口按订单号批量查询供应商退款单 APIResponse
alibaba.wdk.supplier.refund.get

五道口按订单号批量查询供应商退款单
*/
type AlibabaWdkSupplierRefundGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSupplierRefundGetResponse `json:"alibaba_wdk_supplier_refund_get_response,omitempty"`
}

type AlibabaWdkSupplierRefundGetResponse struct {

    // result
    Result   *OrderSyncRefundListResult `json:"result,omitempty"`

}
