package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT设备支持语料获取 APIRequest
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

func NewAlibabaIotDeviceCorpusGetRequest() *AlibabaIotDeviceCorpusGetRequest{
    return &AlibabaIotDeviceCorpusGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIotDeviceCorpusGetRequest) GetApiMethodName() string {
    return "alibaba.iot.device.corpus.get"
}

func (r AlibabaIotDeviceCorpusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIotDeviceCorpusGetRequest) SetUserOpenId(userOpenId string) error {
    r.userOpenId = userOpenId
    r.Set("user_open_id", userOpenId)
    return nil
}

func (r AlibabaIotDeviceCorpusGetRequest) GetUserOpenId() string {
    return r.userOpenId
}

func (r *AlibabaIotDeviceCorpusGetRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaIotDeviceCorpusGetRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaIotDeviceCorpusGetRequest) SetDevId(devId string) error {
    r.devId = devId
    r.Set("dev_id", devId)
    return nil
}

func (r AlibabaIotDeviceCorpusGetRequest) GetDevId() string {
    return r.devId
}

