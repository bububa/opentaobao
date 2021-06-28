package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口按订单号批量查询供应商正向订单 APIResponse
alibaba.wdk.supplier.order.get

五道口按订单号批量查询供应商正向订单
*/
type AlibabaWdkSupplierOrderGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSupplierOrderGetResponse `json:"alibaba_wdk_supplier_order_get_response,omitempty"` 
    AlibabaWdkSupplierOrderGetResponse
}

/* model for simplify = false
type AlibabaWdkSupplierOrderGetResponse struct {

    // result
    
    Result  *struct {
        OrderListSyncPagedResult  *OrderListSyncPagedResult `json:"order_list_sync_paged_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkSupplierOrderGetResponse struct {

    // result
    Result   *OrderListSyncPagedResult `json:"result,omitempty"`

}
