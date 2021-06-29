package gameact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备号 API请求
taobao.de.activity.machineid.get

获取机器设备id
*/
type TaobaoDeActivityMachineidGetRequest struct {
    model.Params
}

// 初始化TaobaoDeActivityMachineidGetRequest对象
func NewTaobaoDeActivityMachineidGetRequest() *TaobaoDeActivityMachineidGetRequest{
    return &TaobaoDeActivityMachineidGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDeActivityMachineidGetRequest) GetApiMethodName() string {
    return "taobao.de.activity.machineid.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDeActivityMachineidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
