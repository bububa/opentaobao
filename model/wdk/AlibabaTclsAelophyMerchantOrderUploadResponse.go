package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家订单数据上传 APIResponse
alibaba.tcls.aelophy.merchant.order.upload

商家订单数据上传
*/
type AlibabaTclsAelophyMerchantOrderUploadAPIResponse struct {
    model.CommonResponse
    AlibabaTclsAelophyMerchantOrderUploadResponse
}

type AlibabaTclsAelophyMerchantOrderUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_order_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 根据站点名称查询产品
    
    ApiResult   *AlibabaTclsAelophyMerchantOrderUploadApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
