package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口供应商维度正向订单拉取 APIResponse
alibaba.wdk.supplier.order.list

五道口供应商维度正向订单拉取
*/
type AlibabaWdkSupplierOrderListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_supplier_order_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *OrderListSyncPagedResult `json:"result,omitempty" xml:"