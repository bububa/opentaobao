package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备控制结果 API请求
alibaba.ailabs.aligenie.iot.device.control.result

智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知
*/
type AlibabaAilabsAligenieIotDeviceControlResultRequest struct {
    model.Params
    // 请求token
    _requestToken   string
    // 设备ID
    _deviceId   string
    // 操作类型 1：控制操作 0:查询
    _type   int64
    // 控制成功
    _control   bool
    // 厂商执行返回ackCode
    _ackCode   string
}

// 初始化AlibabaAilabsAligenieIotDeviceControlResultRequest对象
func NewAlibabaAilabsAligenieIotDeviceControlResultRequest() *AlibabaAilabsAligenieIotDeviceControlResultRequest{
    return &AlibabaAilabsAligenieIotDeviceControlResultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.iot.device.control.result"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestToken Setter
// 请求token
func (r *AlibabaAilabsAligenieIotDeviceControlResultRequest) SetRequestToken(_requestToken string) error {
    r._requestToken = _requestToken
    r.Set("request_token", _requestToken)
    return nil
}

// RequestToken Getter
func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetRequestToken() string {
    return r._requestToken
}
// DeviceId Setter
// 设备ID
func (r *AlibabaAilabsAligenieIotDeviceControlResultRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetDeviceId() string {
    return r._deviceId
}
// Type Setter
// 操作类型 1：控制操作 0:查询
func (r *AlibabaAilabsAligenieIotDeviceControlResultRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetType() int64 {
    return r._type
}
// Control Setter
// 控制成功
func (r *AlibabaAilabsAligenieIotDeviceControlResultRequest) SetControl(_control bool) error {
    r._control = _control
    r.Set("control", _control)
    return nil
}

// Control Getter
func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetControl() bool {
    return r._control
}
// AckCode Setter
// 厂商执行返回ackCode
func (r *AlibabaAilabsAligenieIotDeviceControlResultRequest) SetAckCode(_ackCode string) error {
    r._ackCode = _ackCode
    r.Set("ack_code", _ackCode)
    return nil
}

// AckCode Getter
func (r AlibabaAilabsAligenieIotDeviceControlResultRequest) GetAckCode() string {
    return r._ackCode
}
