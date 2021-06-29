package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取TVS设备列表 APIRequest
alibaba.ailabs.tvs.device.list

获取用户所绑定的TVS设备列表
*/
type AlibabaAilabsTvsDeviceListRequest struct {
    model.Params

}

func NewAlibabaAilabsTvsDeviceListRequest() *AlibabaAilabsTvsDeviceListRequest{
    return &AlibabaAilabsTvsDeviceListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTvsDeviceListRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tvs.device.list"
}

func (r AlibabaAilabsTvsDeviceListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


