package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机支付二维链接获取 APIResponse
alibaba.retail.device.payUrl.get

贩卖机支付二维链接获取
*/
type AlibabaRetailDevicePayUrlGetAPIResponse struct {
    model.CommonResponse
    AlibabaRetailDevicePayUrlGetResponse
}

type AlibabaRetailDevicePayUrlGetResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_device_payUrl_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaRetailDevicePayUrlGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
