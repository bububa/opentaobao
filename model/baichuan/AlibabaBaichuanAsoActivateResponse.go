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
    Response *AlibabaBaichuanAsoActivateResponse `json:"alibaba_baichuan_aso_activate_response,omitempty"`
}

type AlibabaBaichuanAsoActivateResponse struct {

    // result
    Result   *AsoActivateDeviceResult `json:"result,omitempty"`

}
