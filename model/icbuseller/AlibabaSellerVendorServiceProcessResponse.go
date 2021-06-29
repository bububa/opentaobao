package icbuseller

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商客户关联信息 APIResponse
alibaba.seller.vendor.service.process

服务商客户关联信息
*/
type AlibabaSellerVendorServiceProcessAPIResponse struct {
    model.CommonResponse
    AlibabaSellerVendorServiceProcessResponse
}

type AlibabaSellerVendorServiceProcessResponse struct {
    XMLName xml.Name `xml:"alibaba_seller_vendor_service_process_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlibabaSellerVendorServiceProcessResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
