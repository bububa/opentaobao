package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取TVS设备列表 API请求
alibaba.ailabs.tvs.device.list

获取用户所绑定的TVS设备列表
*/
type AlibabaAilabsTvsDeviceListRequest struct {
    model.Params
}

// 初始化AlibabaAilabsTvsDeviceListRequest对象
func NewAlibabaAilabsTvsDeviceListRequest() *AlibabaAilabsTvsDeviceListRequest{
    return &AlibabaAilabsTvsDeviceListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTvsDeviceListRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tvs.device.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTvsDeviceListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
