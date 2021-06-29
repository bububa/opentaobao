package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
iot设备列表变化接口 APIRequest
youku.ott.iot.devicelist.change

iot设备列表变化接口
*/
type YoukuOttIotDevicelistChangeRequest struct {
    model.Params

    // 变更信息
    changeInfo   string 

}

func NewYoukuOttIotDevicelistChangeRequest() *YoukuOttIotDevicelistChangeRequest{
    return &YoukuOttIotDevicelistChangeRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttIotDevicelistChangeRequest) GetApiMethodName() string {
    return "youku.ott.iot.devicelist.change"
}

func (r YoukuOttIotDevicelistChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttIotDevicelistChangeRequest) SetChangeInfo(changeInfo string) error {
    r.changeInfo = changeInfo
    r.Set("change_info", changeInfo)
    return nil
}

func (r YoukuOttIotDevicelistChangeRequest) GetChangeInfo() string {
    return r.changeInfo
}

