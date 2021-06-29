package icbuseller

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国际站服务市场订单列表接口 APIResponse
alibaba.seller.vendor.order.list

返回服务商在服务市场的客户订单
*/
type AlibabaSellerVendorOrderListAPIResponse struct {
    model.CommonResponse
    AlibabaSellerVendorOrderListResponse
}

type AlibabaSellerVendorOrderListResponse struct {
    XMLName xml.Name `xml:"alibaba_seller_vendor_order_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回
    
    Result   *AlibabaSellerVendorOrderListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
