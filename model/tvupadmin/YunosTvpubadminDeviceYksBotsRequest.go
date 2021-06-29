package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备列表 APIRequest
yunos.tvpubadmin.device.yks.bots

获取设备列表
*/
type YunosTvpubadminDeviceYksBotsRequest struct {
    model.Params

}

func NewYunosTvpubadminDeviceYksBotsRequest() *YunosTvpubadminDeviceYksBotsRequest{
    return &YunosTvpubadminDeviceYksBotsRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceYksBotsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.bots"
}

func (r YunosTvpubadminDeviceYksBotsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


