package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备控制结果 APIRequest
alibaba.ailabs.aligenie.iot.device.control.result

智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知
*/
type AlibabaAilabsAligenieIotDeviceControlResultRequest struct {
    model.Params

    // 请求token
    requestToken   string 

    // 设备ID
    deviceId   string 

    // 操作类型 1：控制操作 0:查询
    type   int64 

    // 控制成功
    control   bool 

    // 厂商执行返回ackCode
    ackCode   string 

}

func NewAlibabaAilabsAligenieIotDeviceControlResultRequest() *AlibabaAilabsAligenieIotDeviceControlResultRequest{
    return &AlibabaAilabsAligenieIotDeviceControlResultRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.iot.device.control.result"
}

func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieIotDeviceControlResultRequest) SetRequestToken(requestToken string) error {
    r.requestToken = requestToken
    r.Set("request_token", requestToken)
    return nil
}

func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetRequestToken() string {
    return r.requestToken
}

func (r *AlibabaAilabsAligenieIotDeviceControlResultRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *AlibabaAilabsAligenieIotDeviceControlResultRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetType() int64 {
    return r.type
}

func (r *AlibabaAilabsAligenieIotDeviceControlResultRequest) SetControl(control bool) error {
    r.control = control
    r.Set("control", control)
    return nil
}

func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetControl() bool {
    return r.control
}

func (r *AlibabaAilabsAligenieIotDeviceControlResultRequest) SetAckCode(ackCode string) error {
    r.ackCode = ackCode
    r.Set("ack_code", ackCode)
    return nil
}

func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetAckCode() string {
    return r.ackCode
}

