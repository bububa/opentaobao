package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备id查询设备是否在线 APIRequest
yunos.service.cmns.coa.device.isonline

根据设备id查询设备是否在线
*/
type YunosServiceCmnsCoaDeviceIsonlineRequest struct {
    model.Params

    // 设备id类型，取值"uuid"或者"imei"或者"deviceToken"
    type   string 

    // 对应的设备id值
    value   string 

}

func NewYunosServiceCmnsCoaDeviceIsonlineRequest() *YunosServiceCmnsCoaDeviceIsonlineRequest{
    return &YunosServiceCmnsCoaDeviceIsonlineRequest{
        Params: model.NewParams(),
    }
}

func (r YunosServiceCmnsCoaDeviceIsonlineRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.device.isonline"
}

func (r YunosServiceCmnsCoaDeviceIsonlineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosServiceCmnsCoaDeviceIsonlineRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r YunosServiceCmnsCoaDeviceIsonlineRequest) GetType() string {
    return r.type
}

func (r *YunosServiceCmnsCoaDeviceIsonlineRequest) SetValue(value string) error {
    r.value = value
    r.Set("value", value)
    return nil
}

func (r YunosServiceCmnsCoaDeviceIsonlineRequest) GetValue() string {
    return r.value
}

