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
    Response *AlibabaRetailDevicePayUrlGetResponse `json:"alibaba_retail_device_payUrl_get_response,omitempty"`
}

type AlibabaRetailDevicePayUrlGetResponse struct {

    // result
    Result   *AlibabaRetailDevicePayUrlGetResult `json:"result,omitempty"`

}
