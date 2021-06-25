package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备详情查询 APIRequest
yunos.service.cmns.coa.device.get

第三方应用开发者调用此接口查询设备详情
*/
type YunosServiceCmnsCoaDeviceGetRequest struct {
    model.Params

    // 设备id类型,可以是uuid,imei,deviceToken,kp
    type   string 

    // 设备id
    value   string 

}

func NewYunosServiceCmnsCoaDeviceGetRequest() *YunosServiceCmnsCoaDeviceGetRequest{
    return &YunosServiceCmnsCoaDeviceGetRequest{
        Params: model.NewParams(),
    }
}

func (r YunosServiceCmnsCoaDeviceGetRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.device.get"
}

func (r YunosServiceCmnsCoaDeviceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosServiceCmnsCoaDeviceGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r YunosServiceCmnsCoaDeviceGetRequest) GetType() string {
    return r.type
}

func (r *YunosServiceCmnsCoaDeviceGetRequest) SetValue(value string) error {
    r.value = value
    r.Set("value", value)
    return nil
}

func (r YunosServiceCmnsCoaDeviceGetRequest) GetValue() string {
    return r.value
}

