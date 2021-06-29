package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
iot设备列表变化接口 API请求
youku.ott.iot.devicelist.change

iot设备列表变化接口
*/
type YoukuOttIotDevicelistChangeRequest struct {
    model.Params
    // 变更信息
    _changeInfo   string
}

// 初始化YoukuOttIotDevicelistChangeRequest对象
func NewYoukuOttIotDevicelistChangeRequest() *YoukuOttIotDevicelistChangeRequest{
    return &YoukuOttIotDevicelistChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttIotDevicelistChangeRequest) GetApiMethodName() string {
    return "youku.ott.iot.devicelist.change"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttIotDevicelistChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChangeInfo Setter
// 变更信息
func (r *YoukuOttIotDevicelistChangeRequest) SetChangeInfo(_changeInfo string) error {
    r._changeInfo = _changeInfo
    r.Set("change_info", _changeInfo)
    return nil
}

// ChangeInfo Getter
func (r YoukuOttIotDevicelistChangeRequest) GetChangeInfo() string {
    return r._changeInfo
}
