package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详情 API请求
alibaba.ailabs.tmallgenie.auth.device.get

通过此接口获取设备详情
*/
type AlibabaAilabsTmallgenieAuthDeviceGetRequest struct {
    model.Params
    // 客户id
    clientId   string
    // 用户开放id
    userOpenId   string
    // 设备uuid
    uuid   string
}

// 初始化AlibabaAilabsTmallgenieAuthDeviceGetRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceGetRequest() *AlibabaAilabsTmallgenieAuthDeviceGetRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClientId Setter
// 客户id
func (r *AlibabaAilabsTmallgenieAuthDeviceGetRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetClientId() string {
    return r.clientId
}
// UserOpenId Setter
// 用户开放id
func (r *AlibabaAilabsTmallgenieAuthDeviceGetRequest) SetUserOpenId(userOpenId string) error {
    r.userOpenId = userOpenId
    r.Set("user_open_id", userOpenId)
    return nil
}

// UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetUserOpenId() string {
    return r.userOpenId
}
// Uuid Setter
// 设备uuid
func (r *AlibabaAilabsTmallgenieAuthDeviceGetRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetUuid() string {
    return r.uuid
}
