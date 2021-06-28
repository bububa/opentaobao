package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口按订单号批量查询供应商正向订单 APIResponse
alibaba.wdk.supplier.order.get

五道口按订单号批量查询供应商正向订单
*/
type AlibabaWdkSupplierOrderGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSupplierOrderGetResponse
}

type AlibabaWdkSupplierOrderGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_supplier_order_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *OrderListSyncPagedResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
