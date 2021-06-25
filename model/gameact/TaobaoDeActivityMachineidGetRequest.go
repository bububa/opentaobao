package gameact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备号 APIRequest
taobao.de.activity.machineid.get

获取机器设备id
*/
type TaobaoDeActivityMachineidGetRequest struct {
    model.Params

}

func NewTaobaoDeActivityMachineidGetRequest() *TaobaoDeActivityMachineidGetRequest{
    return &TaobaoDeActivityMachineidGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDeActivityMachineidGetRequest) GetApiMethodName() string {
    return "taobao.de.activity.machineid.get"
}

func (r TaobaoDeActivityMachineidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


