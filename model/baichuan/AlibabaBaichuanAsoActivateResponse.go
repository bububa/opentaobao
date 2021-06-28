package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设备安装活动激活 APIResponse
alibaba.baichuan.aso.activate

设备安装活动激活
*/
type AlibabaBaichuanAsoActivateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaBaichuanAsoActivateResponse `json:"alibaba_baichuan_aso_activate_response,omitempty"` 
    AlibabaBaichuanAsoActivateResponse
}

/* model for simplify = false
type AlibabaBaichuanAsoActivateResponse struct {

    // result
    
    Result  *struct {
        AsoActivateDeviceResult  *AsoActivateDeviceResult `json:"aso_activate_device_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaBaichuanAsoActivateResponse struct {

    // result
    Result   *AsoActivateDeviceResult `json:"result,omitempty"`

}
