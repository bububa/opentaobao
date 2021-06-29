package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
iot设备状态变化通知接口 API请求
youku.ott.iot.status.push

ott iot设备状态通知
*/
type YoukuOttIotStatusPushRequest struct {
    model.Params
    // 变更信息
    changeInfo   string
}

// 初始化YoukuOttIotStatusPushRequest对象
func NewYoukuOttIotStatusPushRequest() *YoukuOttIotStatusPushRequest{
    return &YoukuOttIotStatusPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttIotStatusPushRequest) GetApiMethodName() string {
    return "youku.ott.iot.status.push"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttIotStatusPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChangeInfo Setter
// 变更信息
func (r *YoukuOttIotStatusPushRequest) SetChangeInfo(changeInfo string) error {
    r.changeInfo = changeInfo
    r.Set("change_info", changeInfo)
    return nil
}

// ChangeInfo Getter
func (r YoukuOttIotStatusPushRequest) GetChangeInfo() string {
    return r.changeInfo
}
