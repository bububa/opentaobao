package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口按供应商拉取退款单 APIResponse
alibaba.wdk.supplier.refund.list

五道口按供应商拉取退款单
*/
type AlibabaWdkSupplierRefundListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_supplier_refund_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *OrderSyncRefundListResult `json:"result,omitempty" xml:"