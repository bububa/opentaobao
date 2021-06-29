package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT设备支持语料获取 API请求
alibaba.iot.device.corpus.get

ISV通过该接口获取天猫精灵IoT设备支持控制或查询的语料
*/
type AlibabaIotDeviceCorpusGetRequest struct {
    model.Params
    // 天猫精灵开放用户id
    userOpenId   string
    // 天猫精灵开放的client id
    clientId   string
    // iot设备id
    devId   string
}

// 初始化AlibabaIotDeviceCorpusGetRequest对象
func NewAlibabaIotDeviceCorpusGetRequest() *AlibabaIotDeviceCorpusGetRequest{
    return &AlibabaIotDeviceCorpusGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIotDeviceCorpusGetRequest) GetApiMethodName() string {
    return "alibaba.iot.device.corpus.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIotDeviceCorpusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserOpenId Setter
// 天猫精灵开放用户id
func (r *AlibabaIotDeviceCorpusGetRequest) SetUserOpenId(userOpenId string) error {
    r.userOpenId = userOpenId
    r.Set("user_open_id", userOpenId)
    return nil
}

// UserOpenId Getter
func (r AlibabaIotDeviceCorpusGetRequest) GetUserOpenId() string {
    return r.userOpenId
}
// ClientId Setter
// 天猫精灵开放的client id
func (r *AlibabaIotDeviceCorpusGetRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

// ClientId Getter
func (r AlibabaIotDeviceCorpusGetRequest) GetClientId() string {
    return r.clientId
}
// DevId Setter
// iot设备id
func (r *AlibabaIotDeviceCorpusGetRequest) SetDevId(devId string) error {
    r.devId = devId
    r.Set("dev_id", devId)
    return nil
}

// DevId Getter
func (r AlibabaIotDeviceCorpusGetRequest) GetDevId() string {
    return r.devId
}
