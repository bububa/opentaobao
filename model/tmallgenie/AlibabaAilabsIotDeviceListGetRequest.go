package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取iot设备列表 API请求
alibaba.ailabs.iot.device.list.get

通过此接口获取用户名下的iot设备列表
*/
type AlibabaAilabsIotDeviceListGetRequest struct {
    model.Params
    // 用户open id
    _userOpenId   string
    // client id
    _clientId   string
}

// 初始化AlibabaAilabsIotDeviceListGetRequest对象
func NewAlibabaAilabsIotDeviceListGetRequest() *AlibabaAilabsIotDeviceListGetRequest{
    return &AlibabaAilabsIotDeviceListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceListGetRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.device.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserOpenId Setter
// 用户open id
func (r *AlibabaAilabsIotDeviceListGetRequest) SetUserOpenId(_userOpenId string) error {
    r._userOpenId = _userOpenId
    r.Set("user_open_id", _userOpenId)
    return nil
}

// UserOpenId Getter
func (r AlibabaAilabsIotDeviceListGetRequest) GetUserOpenId() string {
    return r._userOpenId
}
// ClientId Setter
// client id
func (r *AlibabaAilabsIotDeviceListGetRequest) SetClientId(_clientId string) error {
    r._clientId = _clientId
    r.Set("client_id", _clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabsIotDeviceListGetRequest) GetClientId() string {
    return r._clientId
}
