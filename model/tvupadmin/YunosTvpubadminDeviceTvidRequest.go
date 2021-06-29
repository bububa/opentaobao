package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询终端信息 APIRequest
yunos.tvpubadmin.device.tvid

通过UUID查询终端信息
*/
type YunosTvpubadminDeviceTvidRequest struct {
    model.Params

    // 设备的UUID
    uuid   string 

}

func NewYunosTvpubadminDeviceTvidRequest() *YunosTvpubadminDeviceTvidRequest{
    return &YunosTvpubadminDeviceTvidRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceTvidRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.tvid"
}

func (r YunosTvpubadminDeviceTvidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceTvidRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r YunosTvpubadminDeviceTvidRequest) GetUuid() string {
    return r.uuid
}

