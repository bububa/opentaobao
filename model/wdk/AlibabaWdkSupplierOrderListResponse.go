package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口供应商维度正向订单拉取 APIResponse
alibaba.wdk.supplier.order.list

五道口供应商维度正向订单拉取
*/
type AlibabaWdkSupplierOrderListAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSupplierOrderListResponse `json:"alibaba_wdk_supplier_order_list_response,omitempty"`
}

type AlibabaWdkSupplierOrderListResponse struct {

    // result
    Result   *OrderListSyncPagedResult `json:"result,omitempty"`

}
