package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机支付二维链接获取 APIResponse
alibaba.retail.device.payUrl.get

贩卖机支付二维链接获取
*/
type AlibabaRetailDevicePayUrlGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaRetailDevicePayUrlGetResponse `json:"alibaba_retail_device_payUrl_get_response,omitempty"` 
    AlibabaRetailDevicePayUrlGetResponse
}

/* model for simplify = false
type AlibabaRetailDevicePayUrlGetResponse struct {

    // result
    
    Result  *struct {
        AlibabaRetailDevicePayUrlGetResult  *AlibabaRetailDevicePayUrlGetResult `json:"alibaba_retail_device_pay_url_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaRetailDevicePayUrlGetResponse struct {

    // result
    Result   *AlibabaRetailDevicePayUrlGetResult `json:"result,omitempty"`

}
