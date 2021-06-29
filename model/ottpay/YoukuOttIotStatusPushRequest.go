package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
iot设备状态变化通知接口 APIRequest
youku.ott.iot.status.push

ott iot设备状态通知
*/
type YoukuOttIotStatusPushRequest struct {
    model.Params

    // 变更信息
    changeInfo   string 

}

func NewYoukuOttIotStatusPushRequest() *YoukuOttIotStatusPushRequest{
    return &YoukuOttIotStatusPushRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttIotStatusPushRequest) GetApiMethodName() string {
    return "youku.ott.iot.status.push"
}

func (r YoukuOttIotStatusPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttIotStatusPushRequest) SetChangeInfo(changeInfo string) error {
    r.changeInfo = changeInfo
    r.Set("change_info", changeInfo)
    return nil
}

func (r YoukuOttIotStatusPushRequest) GetChangeInfo() string {
    return r.changeInfo
}

