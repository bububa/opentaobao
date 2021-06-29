package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵IoT异步控制回调接口 APIRequest
alibaba.ailabs.iot.device.control.notify

用于天猫精灵IoT云云接入控制结果的异步回调
*/
type AlibabaAilabsIotDeviceControlNotifyRequest struct {
    model.Params

    // 入参
    notifyControlParams   *NotifyVehicleControlParams 

}

func NewAlibabaAilabsIotDeviceControlNotifyRequest() *AlibabaAilabsIotDeviceControlNotifyRequest{
    return &AlibabaAilabsIotDeviceControlNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsIotDeviceControlNotifyRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.device.control.notify"
}

func (r AlibabaAilabsIotDeviceControlNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsIotDeviceControlNotifyRequest) SetNotifyControlParams(notifyControlParams *NotifyVehicleControlParams) error {
    r.notifyControlParams = notifyControlParams
    r.Set("notify_control_params", notifyControlParams)
    return nil
}

func (r AlibabaAilabsIotDeviceControlNotifyRequest) GetNotifyControlParams() *NotifyVehicleControlParams {
    return r.notifyControlParams
}

